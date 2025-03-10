import user from '~/assets/user.svg';
import { Link, Outlet } from 'react-router';
import logo from '~/assets/logo.png';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '~/components/ui/dropdown-menu';
import { NavbarLink } from '../components/navbar-link';
import { useAuth } from '~/hooks/useAuth';

export default function Layout() {
  const { isLoggedIn } = useAuth();

  return (
    <div className="flex flex-col">
      <div className="flex flex-1 justify-center bg-primary">
        <div className="flex w-5/6 justify-between">
          <Link to="/" className="flex items-center gap-2">
            <img className="h-[90px]" src={logo} />
            <h1 className="text-white text-2xl font-bold">NOAH ENGLISH</h1>
          </Link>
          <div className="flex items-center gap-3">
            {/* <NavbarLink to="/">Full Test</NavbarLink> */}
            {isLoggedIn ? (
              <>
                <NavbarLink to="/">Home</NavbarLink>
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
              </>
            ) : (
              <>
                <NavbarLink to="/login">Login</NavbarLink>
                <NavbarLink className="bg-[#EF1B31]" to="/register">
                  Sign Up
                </NavbarLink>
              </>
            )}
          </div>
        </div>
      </div>
      <Outlet />
    </div>
  );
}
