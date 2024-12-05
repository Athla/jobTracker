import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { useState } from "react";
import { useForm } from "react-hook-form";

interface JobFormData {
  name: string;
  source: string;
  description: string;
  created_at: string;
}

const CreateJob: React.FC = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);

  const form = useForm<JobFormData>({
    defaultValues: {
      name: "",
      source: "",
      description: "",
      created_at: (() => {
        const today = new Date();
        today.setHours(12, 0, 0, 0); // Set to 12:00:00
        return today.toISOString().slice(0, 16); // Format as YYYY-MM-DDTHH:mm
      })(),
    },
  });

  const onSubmit = async (data: JobFormData) => {
    setIsSubmitting(true);

    const newJob = {
      ...data,
      created_at: new Date(data.created_at).toISOString(),
    };

    try {
      const response = await fetch("http://localhost:8080/api/jobs/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newJob),
      });

      if (!response.ok) {
        throw new Error("Failed to create job.");
      }

      const responseData = await response.json();
      console.log(responseData);

      form.reset();
      setIsOpen(false);
    } catch (error) {
      console.error("Error", error);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button className="bg-primary hover:bg-primary/90 text-primary-foreground">
          Create Job
        </Button>
      </DialogTrigger>
      <DialogContent className="dialog-content sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Create New Job</DialogTitle>
          <DialogDescription>
            Fill out the form below to create a new job entry.
          </DialogDescription>
          <Button
            onClick={() => setIsOpen(false)}
            className="absolute top-1 right-2 p-1 text-sm"
          >
            Close
          </Button>
        </DialogHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Job Name</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter job name" {...field} />
                  </FormControl>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="source"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Source</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter job source" {...field} />
                  </FormControl>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Description</FormLabel>
                  <FormControl>
                    <Textarea placeholder="Enter job description" {...field} />
                  </FormControl>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="created_at"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Created Date</FormLabel>
                  <FormControl>
                    <Input type="datetime-local" {...field} />
                  </FormControl>
                </FormItem>
              )}
            />
            <Button type="submit" disabled={isSubmitting} className="w-full">
              {isSubmitting ? "Creating..." : "Create Job"}
            </Button>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};

export default CreateJob;
