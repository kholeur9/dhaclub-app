import type { Todo } from "../types/todo"


function TodoItem({todo}:{todo: Todo}) {
    if (todo) {
        alert("Okay")
    }
    return(
        <>
            <div>
                <p className="desc">{todo.description}</p>
            </div>
        </>
    )
}

export default TodoItem