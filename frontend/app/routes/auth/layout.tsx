import { Outlet } from 'react-router';
import asideBanner from '~/assets/aside-banner.png';

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
