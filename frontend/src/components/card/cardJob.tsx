import React from 'react';
import './cardJob.css'
// Create job representation as a interface
// Create a for each, create a smal ui
// Create a job card, detailed and small visions

export interface Job {
  Id: string,
  Name: string,
  Source: string,
  Description: string,
  CreatedAt: string,
}

interface JobCardProps {
  job: Job;
}

const JobCard: React.FC<JobCardProps> = ({ job }) => {
  console.log(job);

  return (
    <div className='job-card'>
    <h1>{job.Name}</h1>
    <p>{job.Description}</p>
    <small>Source: {job.Source}</small>
    <small>Created At: {job.CreatedAt} </small>
    </div>
  )
}

export default JobCard;