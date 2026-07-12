//import { useState } from "react";
import {TodosList} from "../components/TodosList";
import type { Todo } from "../types/todo";

function Dashboard() {
  const todo1: Todo = {
    id: "1",
    description: "MY FIRST TODO",
    completed: false,
  }
  const todos: Todo[] = [todo1];
  return (
    <>
      <h1>Dashboard MIRANDE</h1>
      <TodosList list={todos} />
    </>
  );
}

export default Dashboard;
