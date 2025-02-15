import { Link } from 'react-router';
import { Button } from '~/components/ui/button';
import { Input } from '~/components/ui/input';
import { Label } from '~/components/ui/label';
import google from '~/assets/google.svg';
import { ChevronLeft } from 'lucide-react';

export default function Login() {
  return (
    <div className="flex flex-col gap-12 p-8">
      <Link to="/" className="flex items-center gap-1.5 text-secondary">
        <ChevronLeft /> Back
      </Link>
      <div className="flex flex-col gap-14 w-2/3 mx-auto">
        <div className="flex flex-col gap-6">
          <h1 className="text-4xl font-bold">Login</h1>
          <h2 className="text-lg text-secondary">Welcome back to Noah English!</h2>
        </div>
        <div className="flex flex-col gap-10">
          <form className="flex flex-col gap-5">
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="email" className="text-secondary">
                Email
              </Label>
              <Input id="email" placeholder="Enter your email" className="h-14" />
            </div>
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="password" className="text-secondary">
                Password
              </Label>
              <Input id="password" placeholder="Enter your password" className="h-14" />
            </div>
            <Button className="h-14 text-md">Sign in</Button>
          </form>
          <Button className="h-14 text-md" variant="outline">
            Sign in with Google <img src={google} />
          </Button>
        </div>
      </div>
    </div>
  );
}
