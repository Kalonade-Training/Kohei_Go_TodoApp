package handlers

import (
	"net/http"
	"time"
	value_object "goTodoApp/domain/value-object"
	req "goTodoApp/interface-adapter/dto/request"
	res "goTodoApp/interface-adapter/dto/response"
	"goTodoApp/common"
	"goTodoApp/application/user/todo"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"	
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	createTodoUC   *todo.CreateTodoUseCase
	findTodoByIDUC *todo.FindTodoByIDUseCase
	findByUserIDWithFiltersUC *todo.FindByUserIDWithFiltersUseCase
	updateTodoUC   *todo.UpdateTodoUseCase
	deleteTodoUC  *todo.DeleteTodoUseCase
	duplicateTodoUC *todo.DuplicateTodoUseCase
}

func NewTodoController(
	create *todo.CreateTodoUseCase,
	findTodoByID *todo.FindTodoByIDUseCase,
	findByUserIDWithFilters *todo.FindByUserIDWithFiltersUseCase,
	update *todo.UpdateTodoUseCase,
	deleteUC *todo.DeleteTodoUseCase,
	duplicate *todo.DuplicateTodoUseCase,
) *TodoController {
	return &TodoController{
		createTodoUC:    create,
		findTodoByIDUC:  findTodoByID,
		findByUserIDWithFiltersUC:  findByUserIDWithFilters,
		updateTodoUC:    update,
		deleteTodoUC:    deleteUC,
		duplicateTodoUC: duplicate,
	}
}

// Create ハンドラー例
func (tc *TodoController) Create(c *gin.Context) {
	// ① JWTトークンからユーザーIDを取得
	authUserID, err := common.GetAuthUserID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID in token"})
			return
		}

	// ② リクエストを構造体にバインド
	var todoDTO req.Todo
	if err := c.ShouldBindJSON(&todoDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ③ バリューオブジェクトの生成
	// Title
	title, err := value_object.NewTitle(todoDTO.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Description
	var description *value_object.Description
	if todoDTO.Description != nil {
		description, err = value_object.NewDescription(*todoDTO.Description)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// DueDate
	var dueDate *value_object.DueDate
	if todoDTO.DueDate != nil {
		dueDate, err = value_object.NewDueDate(*todoDTO.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	//Status
	var status *value_object.Status
	if todoDTO.Status != nil {
		status, err = value_object.NewStatus(*todoDTO.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	// ④ 入力DTO → ユースケース入力にマッピング
	input := todo.CreateTodoInput{
		UserID:      authUserID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      *status,
	}

	// ⑤ ユースケースを実行
	createdTodo, err := tc.createTodoUC.Execute(input)
	if err != nil {
		handleError(c, err)
		return
	}

	// ⑥ エンティティ → DTO に変換
	todoResponse := entityToDTO(createdTodo)

	// ⑦ レスポンスを返す
	c.JSON(http.StatusCreated, todoResponse)
	return
}

func (tc *TodoController) FindByUserIDWithFilters(c *gin.Context) {
	// JWTからUserIDを取得
	authUserID, err := common.GetAuthUserID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
			return
		}
	//requestパラメータを取得し、filtersDTOにバインド
	var filtersDTO req.TodoFilters
	if err := c.ShouldBindQuery(&filtersDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// dto → repository用フィルタに変換
	var dueDateFrom, dueDateTo *value_object.DueDate
	if filtersDTO.DueDateFrom != nil && *filtersDTO.DueDateFrom != ""{
		dueDateFrom, err = value_object.NewDueDate(*filtersDTO.DueDateFrom)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_date_from"})
			return
		}
	}	

	if filtersDTO.DueDateTo != nil && *filtersDTO.DueDateTo != ""	{
		dueDateTo, err = value_object.NewDueDate(*filtersDTO.DueDateTo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_date_from"})
			return
		}
	}
	var dueDateFromTime, dueDateToTime *time.Time
	if dueDateFrom != nil {
		val := dueDateFrom.Value()
		dueDateFromTime = &val
	}
	if dueDateTo != nil {
		val := dueDateTo.Value()
		dueDateToTime = &val
	}

	domainFilters := repositories.TodoFilters{
		Title:       filtersDTO.Title,
		Description: filtersDTO.Description,
		DueDateFrom: dueDateFromTime,
		DueDateTo:   dueDateToTime,
		Status:      filtersDTO.Status,
	}

	todos, err := tc.findByUserIDWithFiltersUC.Execute(authUserID, domainFilters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//レスポンス用のDTOに変換
	todoResponses := make([]res.TodoResponse, 0, len(todos))
	for _, todo := range todos {
		todoResponses = append(todoResponses, entityToDTO(todo))
	}
	c.JSON(http.StatusOK, todoResponses)
	return
}

func (tc *TodoController) FindTodoByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo ID is required"})
		return
	}

	// JWTからUserIDを取得
	authUserID, err := common.GetAuthUserID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
			return
		}

	// TodoIDのバリデーションとVO変換
	todoID, err := value_object.FromStringTodoID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	// ユースケース実行
	todo, err := tc.findTodoByIDUC.Execute(todoID, authUserID)
	if err != nil {
		handleError(c, err)
		return
	}

	// DTOに変換して返却
	todoResponse := entityToDTO(todo)

	c.JSON(http.StatusOK, todoResponse)
}

func (tc *TodoController) Update(c *gin.Context) {
	id := c.Param("id")

	// JWTトークンからユーザーIDを取得
	authUserID, err := common.GetAuthUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID in token"})
		return
	}

	// リクエストBodyをDTO構造体にバインド
	var userInput req.UpdateTodoRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	// 値ごとのバリデーションとVO生成
	var title *value_object.Title
	if userInput.Title != nil {
    t, err := value_object.NewTitle(*userInput.Title)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    title = &t
	}

	var description *value_object.Description
	if userInput.Description == nil ||  (userInput.Description != nil && *userInput.Description == "") {
    description = nil
	} else if *userInput.Description == "" {
    description = nil
		} else {
    description, err = value_object.NewDescription(*userInput.Description)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	}

	var dueDate *value_object.DueDate
	if userInput.DueDate == nil || (userInput.DueDate != nil && *userInput.DueDate == "") {
    // due_date が null または空文字の場合は nil（削除扱い）
    dueDate = nil
	} else {
    // それ以外はパースしてバリデーション
    dueDate, err = value_object.NewDueDate(*userInput.DueDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	}
	var status *value_object.Status
	if userInput.Status == nil || (userInput.Status != nil && *userInput.Status == "") {
		status = nil
	} else {
		 // それ以外はパースしてバリデーション
    status, err = value_object.NewStatus(*userInput.Status)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
		}
	}

	var completedAt *value_object.CompletedAt
	if userInput.CompletedAt != nil {
		parsed, err := time.Parse("2006-01-02", *userInput.CompletedAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid completed_at format. Use YYYY-MM-DD"})
        return
		}
		completedAt = value_object.NewCompletedAt(parsed)
	}

	// ユースケース実行
	input := todo.UpdateTodoInput{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
		CompletedAt: completedAt,
	}
	todoID, err := value_object.FromStringTodoID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := tc.updateTodoUC.Execute(todoID, authUserID, input)
	if err != nil {
		handleError(c, err)
		return
	}
	
	// エンティティ → DTO に変換して返却
	todoResponse := entityToDTO(updatedTodo)

	c.JSON(http.StatusOK, todoResponse)
	return
}

func (tc *TodoController) Delete(c *gin.Context) {
	id := c.Param("id")
	// JWTトークンからユーザーIDを取得
	authUserID, err := common.GetAuthUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID in token"})
		return
	}

	// TodoIDの検証
	todoID, err := value_object.FromStringTodoID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}
	// ユースケース実行
	if err := tc.deleteTodoUC.Execute(todoID, authUserID); err != nil {
		handleError(c, err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Todo successfully deleted"} ) 
	return
}

func (tc *TodoController) Duplicate(c *gin.Context) {
	id := c.Param("id")

	// JWTトークンからユーザーIDを取得
	authUserID, err := common.GetAuthUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID in token"})
		return
	}

	// TodoIDの検証
	todoID, err := value_object.FromStringTodoID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	// ユースケース実行
	duplicatedTodo, err := tc.duplicateTodoUC.Execute(todoID, authUserID)
	if err != nil {
		handleError(c, err)
		return
	}

	// DTOに変換して返却
	todoResponse := entityToDTO(duplicatedTodo)
	
	c.JSON(http.StatusCreated, todoResponse) // 201 Created
}


// entityToDTOはエンティティからレスポンス用のDTOに変換する
func entityToDTO(todo *entities.Todo) res.TodoResponse {
	// Descriptionのnilチェック
	var description *string
	if todo.Description() != nil {
		val := todo.Description().Value()
		description = &val
	}
	// DueDateのnilチェック
	var dueDate *string
	if todo.DueDate() != nil {
		val := todo.DueDate().Value().Format("2006-01-02")
		dueDate = &val
	}
	// CompletedAtのnilチェック
	var completedAt *time.Time
	if todo.CompletedAt() != nil {
		val := todo.CompletedAt().Value()
		completedAt = &val
	}
	// Statusのnilチェック
	var status *string
	if todo.Status().Value() != "" {
    val := todo.Status().Value()
    status = &val
	}

	return res.TodoResponse{
		ID:          todo.ID().Value(),
		UserID:      todo.UserID().Value(),
		Title:       todo.Title().Value(),
		Description: description,
		DueDate:     dueDate,
		Status:      status,
		CompletedAt: completedAt,
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
}

func handleError(c *gin.Context, err error) {
	switch err.Error() {
	case "todo not found":
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}