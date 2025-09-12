export type CourseInfo = {
    instructor: string;
    code: string;
};

export type ClassData = Record<string, CourseInfo>;

export type DataStructure = {
    "BsCS-4C": ClassData;
    "BsCS-5A": ClassData;
};

export const data: DataStructure = {
    "BsCS-4C": {
        "Database Systems": {
            instructor: "Dr. Shahzad Latif",
            code: "CSC 2203",
        },
        "Finite Automata Theory and Formal Languages": {
            instructor: "Dr. Shahzad Latif",
            code: "CSC 2204",
        },
        "Linear Algebra": {
            instructor: "Dr. Shahzad Latif",
            code: "CSC 2206",
        },
        "Design and Analysis of Algorithms": {
            instructor: "Dr. Shahzad Latif",
            code: "CSC 3202",
        },
    },
    "BsCS-5A": {
        "Natural Language Processing": {
            instructor: "Saad Irfan khan",
            code: "CSC 4207",
        },
        "Technical and Business Writing": {
            instructor: "Muhammad Qasim",
            code: "ENG 3201",
        },
        "Computer Networks and Data Communications": {
            instructor: "Rana Faisal Hayat",
            code: "CSC 4203",
        },
        "Artificial Intelligence": {
            instructor: "Muhammad Qasim",
            code: "CSC 4204",
        },
        "Web Technologies-I": {
            instructor: "Muhammad Qasim",
            code: "CSC 4205",
        },
        "Android Application Development": {
            instructor: "Zubair Ahmed Chatta",
            code: "CSC 4206",
        },
        "Lab: Computer Networks and Data Communications": {
            instructor: "Muhammad Qasim",
            code: "CSC 4208",
        },
        "Lab: Artificial Intelligence": {
            instructor: "Muhammad Qasim",
            code: "CSC 4209",
        },
    },
};