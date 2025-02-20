import user from '~/assets/user.svg';
import { Link, NavLink, Outlet } from 'react-router';
import logo from '~/assets/logo.png';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '~/components/ui/dropdown-menu';
import { cn } from '~/lib/utils';
import type { ReactNode } from 'react';

interface LinkProps {
  to: string;
  children: ReactNode;
}
export function NavbarLink({ to, children }: LinkProps) {
  return (
    <NavLink
      to={to}
      className={({ isActive }) =>
        cn(
          'text-md font-semibold text-white hover:text-black hover:bg-white px-4 py-2 rounded-full',
          isActive ? 'text-black bg-white' : ''
        )
      }
    >
      {children}
    </NavLink>
  );
}

export default function Layout() {
  return (
    <div className="flex flex-col">
      <div className="flex flex-1 justify-center bg-primary">
        <div className="flex w-5/6 justify-between">
          <Link to="/" className="flex items-center gap-2">
            <img className="h-[90px]" src={logo} />
            <h1 className="text-white text-2xl font-bold">Noah English</h1>
          </Link>
          <div className="flex items-center gap-3">
            <NavbarLink to="/">Full Test</NavbarLink>
            <NavbarLink to="/section">Section Test</NavbarLink>
            <NavbarLink to="/review">Review</NavbarLink>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <button className="font-semibold text-white bg-[#D9D9D9] hover:bg-[#cccbcb] p-3 ml-0.5 rounded-full outline-none">
                  <img className="h-6 w-6" src={user} />
                </button>
              </DropdownMenuTrigger>
              <DropdownMenuContent side="bottom" sideOffset={6} align="end">
                <DropdownMenuItem asChild>
                  <Link to="/logout" className="text-md px-4 py-2 font-medium">
                    Logout
                  </Link>
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </div>
      </div>
      <div className="flex flex-col w-5/6 mx-auto my-4">
        <Outlet />
      </div>
    </div>
  );
}
