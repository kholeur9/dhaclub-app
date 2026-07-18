import { useState } from "react";
import {TodosList} from "@/components/TodosList";
import type { Todo } from "@/types/todo";
import { Button } from "@/components/ui/button";

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
    <div>
      <h1>Dashboard MIRANDE</h1>
      <Button onClick={addTodo}>+ add a todo</Button>
      <TodosList list={todos} />
    </div>
  );
}

export default Dashboard;
