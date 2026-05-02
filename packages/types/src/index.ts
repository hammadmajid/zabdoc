import { z } from "zod";

export const documentSchema = z.object({
    Students: z.array(z.object({
        Name: z.string().optional(),
        RegNo: z.string().optional(),
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

export const scrapedData = z.record(
    z.string(), // course name as key
    z.object({
        attendance: z.object({
            courseName: z.string(),
            instructor: z.string(),
            records: z.array(z.object({
                lecture: z.string(),
                date: z.string(), // "dd/mm" format from Go scraper
                status: z.string(),
            })),
        }),
        marks: z.object({
            entries: z.array(z.object({
                head: z.string(),
                max: z.string(),
                obtained: z.string(),
            })),
        }),
    })
);

export type ScrapedData = z.infer<typeof scrapedData>;
