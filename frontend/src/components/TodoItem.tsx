import type { Todo } from "@/types/todo";
import { Item, ItemContent, ItemDescription, ItemMedia } from "./ui/item";

export function TodoItem({ todo }: { todo: Todo }) {
  const { id, description, completed } = todo;
  return (
    <>
      <Item variant="outline" className="w-md" onClick={() => alert(id)}>
        <ItemMedia>
          {completed ? (
            <span className="">✅</span>
          ) : (
            <span className="">▢</span>
          )}
        </ItemMedia>
        <ItemContent>
          <ItemDescription>{description}</ItemDescription>
        </ItemContent>
      </Item>
    </>
  );
}
