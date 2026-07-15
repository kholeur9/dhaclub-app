import type { Todo } from "@/types/todo"
import {TodoItem} from "@/components/TodoItem"

type TodosListProps = {
    list: Todo[]
}

export function TodosList({list}: TodosListProps) {
    return(
        <>
            <h1>Mes TODOS</h1>
            {list.map(todo => {
                return <TodoItem key={todo.id} todo={todo} />
            })}
        </>
    )
}