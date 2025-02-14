import axios from 'axios';
import type { Route } from './+types/test';

interface Todo {
  userId: number;
  id: number;
  title: string;
  completed: boolean;
}

export function meta({}: Route.MetaArgs) {
  return [{ title: 'Test' }];
}

export async function loader({}: Route.LoaderArgs) {
  const response = await axios.get<Todo>('https://jsonplaceholder.typicode.com/posts/1');
  return response.data;
}

export default function Product({ loaderData }: Route.ComponentProps) {
  const { id, title, completed } = loaderData;
  return (
    <div>
      <p>ID: {id}</p>
      <h1>{title}</h1>
      <p>Completed: {completed ? 'Yes' : 'No'}</p>
    </div>
  );
}
