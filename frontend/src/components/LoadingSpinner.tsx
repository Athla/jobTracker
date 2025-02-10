import { cn } from "@/lib/utils";

interface LoadingSpinnerProps {
  size?: "small" | "medium" | "large";
  className?: string;
}

export function LoadingSpinner({
  size = "medium",
  className,
}: LoadingSpinnerProps) {
  return (
    <div
      className={cn(
        "animate-spin rounded-full border-4 border-primary border-t-transparent",
        {
          "h-4 w-4": size === "small",
          "h-8 w-8": size === "medium",
          "h-12 w-12": size === "large",
        },
        className
      )}
    />
  );
}
