import type { Todo } from "../types/todo";

function TodoItem({ todo }: { todo: Todo }) {
  const { id, description, completed } = todo;
  return (
    <>
      <div className="desc" onClick={() => alert(id)}>
        <p className="desc-p">
          {completed ? <span className="desc-c">✅</span> : <span className="desc-c">▢</span>}
          {description}
        </p>
      </div>
    </>
  );
}

export default TodoItem;
