import { Link } from 'react-router';
import { Button } from '~/components/ui/button';
import { Input } from '~/components/ui/input';
import { Label } from '~/components/ui/label';
import google from '~/assets/google.svg';
import { ChevronLeft } from 'lucide-react';

export default function Register() {
  return (
    <div className="flex flex-col gap-12 p-8">
      <Link to="/" className="flex items-center gap-1.5 text-secondary">
        <ChevronLeft /> Back
      </Link>
      <div className="flex flex-col gap-14 w-2/3 mx-auto">
        <div className="flex flex-col gap-6">
          <h1 className="text-4xl font-bold">Sign up</h1>
          <h2 className="text-lg text-secondary">Welcome back to Noah English!</h2>
        </div>
        <div className="flex flex-col gap-10">
          <form className="flex flex-col gap-5">
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="fullname" className="text-secondary">
                Full Name
              </Label>
              <Input id="fullname" placeholder="Enter your name" className="h-14" />
            </div>
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
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="confirm-password" className="text-secondary">
                Confirm Password
              </Label>
              <Input id="confirm-password" placeholder="Enter your password again" className="h-14" />
            </div>
            <Button className="h-14 text-md">Create Account</Button>
          </form>
          <Button className="h-14 text-md" variant="outline">
            Continue with Google <img src={google} />
          </Button>
        </div>
      </div>
    </div>
  );
}
