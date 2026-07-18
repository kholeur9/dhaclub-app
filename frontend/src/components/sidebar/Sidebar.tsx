import {DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuGroup, DropdownMenuLabel} from "@/components/ui/dropdown-menu"
import {Avatar, AvatarImage, AvatarFallback} from "@/components/ui/avatar"
import {Button} from "@/components/ui/button"
import { Plus, SettingsIcon, UserIcon } from "lucide-react"

//import {SidebarItem} from "@/components/sidebar/SidebarItem"

export function Sidebar() {
    return(
        <div className="h-full w-[270px] bg-[#FAFAFA] flex flex-col px-[8px] py-[5px] p-0">
            <DropdownMenu>
                <DropdownMenuTrigger render={
                    <Button variant="ghost" size="sm" className="">
                        <Avatar size="xs">
                            <AvatarImage src="https://github.com/shadcn.png" />
                            <AvatarFallback>YW</AvatarFallback>
                        </Avatar>
                        <span className="text-[14px] font-[500]">Yann Wilfried Moutsinga</span>
                    </Button>
                } />
                <DropdownMenuContent className="bg-white px-1.5 py-1.5 w-[300px] rounded-[8px] shadow-md">
                    <DropdownMenuGroup className="py-0">
                        <DropdownMenuLabel className="text-[12px] py-0">AS: @yannewilfriedm</DropdownMenuLabel>
                        <DropdownMenuItem className="text-[#02F5A1] hover:text-[#02F5A1] gap-1.5 px-2 py-0 text-[14px] rounded-[8px]">
                            <Plus />
                            Ajouter un profil
                        </DropdownMenuItem>
                    </DropdownMenuGroup>
                    <DropdownMenuSeparator className="mx-2" />
                    <DropdownMenuItem className="gap-1.5 px-2 py-0 text-[14px] rounded-[8px]">
                        <UserIcon />
                        Mon compte
                    </DropdownMenuItem>
                    <DropdownMenuItem className="gap-1.5 px-2 py-0 text-[14px] rounded-[8px]">
                        <SettingsIcon />
                        Paramètres
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    )
}