import type { Todo } from "../types/todo"


function TodoItem({todo}:{todo: Todo}) {
    return(
        <>
            <div>
                <p className="desc">{todo.description}</p>
            </div>
        </>
    )
}

export default TodoItem