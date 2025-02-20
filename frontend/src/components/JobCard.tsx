import { Job } from "@/types/job";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { cn } from "@/lib/utils";
import { format } from "date-fns";
import { useJobs } from "@/hooks/use-jobs";
import { TrashIcon } from "lucide-react";
import { Button } from "./ui/button";

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
  const { deleteJob } = useJobs()
  return (
    <Card className="transition-all duration-200 ease-in-out hover:shadow-lg">
      <CardHeader className="p-4">
        <div className="space-y-2">
          <CardTitle className="text-lg font-semibold line-clamp-2">
            {job.name}
          </CardTitle>
          <p className="text-sm font-medium text-muted-foreground">
            {job.company}
          </p>
          <span
            role="status"
            aria-label={`Status: ${job.status}`}
            className={cn(
              "inline-block px-2 py-1 rounded-full text-xs font-medium",
              statusColors[job.status]
            )}
          >
            {job.status.replace(/_/g, " ")}
          </span>
        </div>
      </CardHeader>
      <CardContent className="p-4 pt-0">
        <div className="space-y-2">
          {job.source && (
            <p className="text-sm text-muted-foreground">{job.source}</p>
          )}
          {job.description && (
            <p className="text-sm line-clamp-2">{job.description}</p>
          )}
          <p className="text-xs text-muted-foreground">
            Created: {format(new Date(job.created_at), "MMM d, yyyy")}
          </p>
        </div>
      </CardContent>
      <CardContent>
        <Button variant={"destructive"} size={"icon"} onClick={() => deleteJob(job.id, job.version)}><TrashIcon /> </Button></CardContent>
    </Card>
  );
};

export default JobCard;
