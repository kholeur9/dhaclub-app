type TodoType = {
    id: string
    description: string,
    completed: boolean,
}

export function TodosList({list}:{list: TodoType[]}) {
    return(
        <>
            <h1>Mes TODOS</h1>
            {console.log(list)}
        </>
    )
}