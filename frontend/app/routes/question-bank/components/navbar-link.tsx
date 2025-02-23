import type { ReactNode } from 'react';
import { NavLink } from 'react-router';
import { cn } from '~/lib/utils';

interface NavbarLink {
  to: string;
  children: ReactNode;
}
export function NavbarLink({ to, children }: NavbarLink) {
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
