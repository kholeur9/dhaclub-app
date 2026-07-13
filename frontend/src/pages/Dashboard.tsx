import { useState } from "react";
import {TodosList} from "../components/TodosList";
import type { Todo } from "../types/todo";

function Dashboard() {
  const [todos, setTodos] = useState<Todo[]>([
    {
      id: "1",
      description: "MY FIRST TODO",
      completed: false,
      createdAt: new Date(),
    }
  ])

  const todo2: Todo = {
    id: "2",
    description: "LEARN REACT",
    completed: true,
    createdAt: new Date(),
  }

  const addTodo = () => {
    setTodos((prev) => [...prev, todo2])
  }
  
  return (
    <>
      <h1>Dashboard MIRANDE</h1>
      <button onClick={addTodo}>+ add a todo</button>
      <TodosList list={todos} />
    </>
  );
}

export default Dashboard;
