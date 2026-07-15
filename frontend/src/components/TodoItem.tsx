import type { Todo } from "@/types/todo";

export function TodoItem({ todo }: { todo: Todo }) {
  const { id, description, completed } = todo;
  return (
    <>
      <div className="desc" onClick={() => alert(id)}>
        <p className="desc-p">
          {completed ? <span className="">✅</span> : <span className="">▢</span>}
          {description}
        </p>
      </div>
    </>
  );
}
