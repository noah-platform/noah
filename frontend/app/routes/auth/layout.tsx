import { Outlet } from 'react-router';
import asideBanner from '~/assets/aside-banner.png';
import type { Route } from './+types/layout';

export async function loader({}: Route.LoaderArgs) {
  // console.log('request', request.headers.get('cookie'));
}

export default function Layout() {
  return (
    <div className="grid lg:grid-cols-2 min-h-dvh">
      <div className="hidden lg:flex flex-col justify-center items-center bg-primary">
        <img className="w-full max-w-[700px]" src={asideBanner} />
      </div>
      <div>
        <Outlet />
      </div>
    </div>
  );
}
