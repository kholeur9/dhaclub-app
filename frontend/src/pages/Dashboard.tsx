//import { useState } from "react";
import {TodosList} from "../components/TodosList";
import type { Todo } from "../types/todo";

function Dashboard() {
  const todo1: Todo = {
    id: "1",
    description: "MY FIRST TODO",
    completed: false,
    createdAt: new Date(),
  }
  const todo2: Todo = {
    id: "2",
    description: "LEARN REACT",
    completed: false,
    createdAt: new Date(),
  }
  const todos: Todo[] = [todo1, todo2];
  return (
    <>
      <h1>Dashboard MIRANDE</h1>
      <TodosList list={todos} />
    </>
  );
}

export default Dashboard;
