import { LoadingSpinner } from "./LoadingSpinner";

interface LoadingStateProps {
  message?: string;
}

export function LoadingState({ message = "Loading..." }: LoadingStateProps) {
  return (
    <div className="flex flex-col items-center justify-center space-y-4 p-4">
      <LoadingSpinner size="large" />
      <p className="text-muted-foreground">{message}</p>
    </div>
  );
}
