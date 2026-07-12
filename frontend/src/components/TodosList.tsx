type TodoType = {
    id: string
    description: string,
    completed: boolean,
}

type TodosListProps = {
    list: TodoType[]
}

export function TodosList({list}: TodosListProps) {
    return(
        <>
            <h1>Mes TODOS</h1>
            {console.log(list)}
        </>
    )
}