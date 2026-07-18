import {DropdownMenuItem} from "@/components/ui/dropdown-menu"

export function SidebarItem(children) {
    return(
        <DropdownMenuItem>
            {children}
        </DropdownMenuItem>
    )
}