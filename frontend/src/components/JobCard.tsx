import { Job } from "@/types/job";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { cn } from "@/lib/utils";

interface JobCardProps {
  job: Job;
}

const statusColors: Record<string, string> = {
  WISHLIST: "bg-blue-100 text-blue-800",
  APPLIED: "bg-yellow-100 text-yellow-800",
  PHONE_SCREEN: "bg-purple-100 text-purple-800",
  TECHNICAL_INTERVIEW: "bg-indigo-100 text-indigo-800",
  ONSITE: "bg-pink-100 text-pink-800",
  OFFER: "bg-green-100 text-green-800",
  ACCEPTED: "bg-emerald-100 text-emerald-800",
  REJECTED: "bg-red-100 text-red-800",
  WITHDRAWN: "bg-gray-100 text-gray-800",
};

const JobCard: React.FC<JobCardProps> = ({ job }) => {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  };

  return (
    <Card className="hover:shadow-lg transition-shadow">
      <CardHeader className="p-4">
        <div className="flex justify-between items-start">
          <CardTitle className="text-base font-semibold line-clamp-2">
            {job.name}
          </CardTitle>
          <span
            className={cn(
              "px-2 py-1 rounded-full text-xs font-medium",
              statusColors[job.status],
            )}
          >
            {job.status.replace("_", " ")}
          </span>
        </div>
        <p className="text-sm text-muted-foreground">{job.company}</p>
      </CardHeader>
      <CardContent className="p-4 pt-0">
        <div className="space-y-2">
          <p className="text-sm text-muted-foreground">{job.source}</p>
          {job.description && (
            <p className="text-sm line-clamp-2">{job.description}</p>
          )}
          <p className="text-xs text-muted-foreground">
            Created: {formatDate(job.created_at)}
          </p>
        </div>
      </CardContent>
    </Card>
  );
};

export default JobCard;
