'use client';

import Image from 'next/image';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { Home, ChartLine, WalletMinimal, Tornado, BookText, ArrowDownUp } from "lucide-react";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarFooter,
  SidebarHeader,
} from '@/components/ui/sidebar';

const items = [
  {
    title: 'Home',
    url: '/',
    icon: Home,
  },
  {
    title: 'Accounts',
    url: '/accounts',
    icon: BookText,
  },
  {
    title: 'Records',
    url: '/records',
    icon: Tornado,
  },
  {
    title: 'Assets',
    url: '/assets',
    icon: WalletMinimal,
  },
  {
    title: 'Statistics',
    url: '/statistics',
    icon: ChartLine,
  },
  {
    title: 'Budgets',
    url: '/budgets',
    icon: ArrowDownUp,
  },
]

export function AppSidebar() {
  const pathname = usePathname();

  return (
    <Sidebar>
      <SidebarHeader>
        <Image src="/logo-bgless.png" alt="Vinance Logo" width={128} height={128} />
      </SidebarHeader>
        
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => {
                const isActive = item.url === '/' 
                  ? pathname === '/' || pathname === '/home'
                  : pathname === item.url || pathname.startsWith(`${item.url}/`);
              
              return (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild isActive={isActive}>
                    <Link href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
                );
              })}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter />
    </Sidebar>
  );
}
