import { z } from "zod";

export const documentSchema = z.object({
    Students: z.array(z.object({
        Name: z.string().min(1),
        RegNo: z.string().min(1),
    })),
    Class: z.string().min(1),
    Course: z.string().min(1),
    CourseCode: z.string().min(1),
    Instructor: z.string().min(1),
    DocType: z.string().min(1),
    Number: z.string().min(1),
    Date: z.string().min(1),
    Marks: z.string().min(1),
});

export type DocumentSchema = z.infer<typeof documentSchema>;

export const scrapeInitSchema = z.object({
    username: z.string().min(1),
    password: z.string().min(1),
    campus: z.enum(['isb']),
    semester: z.enum(['fall', 'spring', 'summer']),
});

export type ScrapeInitSchema = z.infer<typeof scrapeInitSchema>;

export const scrapedData = z.object({
    username: z.string(),
    name: z.string(),
    cgpa: z.number(),
    courses: z.array(z.object({
        instructor: z.string(),
        credits: z.number(),
        attendence: z.object({
            lecture: z.number(),
            date: z.date(),
            status: z.string(),
        }),
        marks: z.object({
            head: z.string(),
            total: z.number(),
            obtained: z.number().optional() // optional means "Not entered"
        })
    }))
});

export type ScrapedData = z.infer<typeof scrapedData>;
