import { Button } from '@/components/ui/button';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="text-center">
        <h1 className="text-4xl font-bold mb-4">Welcome to Vinance</h1>
        <Button>Get started</Button>
      </div>
    </main>
  );
}
