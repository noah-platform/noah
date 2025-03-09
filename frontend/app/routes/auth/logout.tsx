import { useNavigate } from 'react-router';
import { ChevronLeft } from 'lucide-react';

import { client } from '~/clients/client';
import { useEffect } from 'react';
import toast from 'react-hot-toast';

export default function Logout() {
  const navigate = useNavigate();
  const { mutateAsync } = client.useMutation('post', '/auth/v1/logout');

  useEffect(() => {
    const logout = async () => {
      try {
        await mutateAsync({});
        toast.dismiss();
        toast.success('Logged out successfully');
        navigate('/');
      } catch {
        toast.error('Something went wrong');
        return;
      }
    };
    logout();
  }, []);

  return (
    <div className="flex flex-col gap-12 p-8">
      <button className="flex items-center gap-1.5 text-secondary" onClick={() => navigate(-1)}>
        <ChevronLeft /> Back
      </button>
      <div className="flex flex-col gap-14 w-7/8 md:w-5/6 lg:w-2/3 mx-auto">
        <div className="flex flex-col gap-6">
          <h1 className="text-4xl font-bold">Logout</h1>
          <h2 className="text-lg text-secondary">Please wait a moment...</h2>
        </div>
      </div>
    </div>
  );
}
