export type CourseInfo = {
    instructor: string;
    code: string;
};

export type ClassData = Record<string, CourseInfo>;

export type DataStructure = {
    "BsCS-4C": ClassData;
    "BsCS-6A": ClassData;
};

export const data: DataStructure = {
    "BsCS-4C": {
        "Database Systems": {
            instructor: "Muhammad Qasim",
            code: "CSC 2203",
        },
        "Finite Automata Theory and Formal Languages": {
            instructor: "Taseer Ul Islam",
            code: "CSC 2204",
        },
        "Linear Algebra": {
            instructor: "Dr. Sajjad Ghouri",
            code: "CSC 2206",
        },
        "Design and Analysis of Algorithms": {
            instructor: "Aniqa",
            code: "CSC 3202",
        },
        "Management Principle": {
            instructor: "Dr. Beenish Ambreen",
            code: "CSC 4603"
        },
        "Lab: Database Systems": {
            instructor: "Mr. Farhan Sami",
            code: "CSCL 2203",
        },
    },
    "BsCS-6A": {
        "Natural Language Processing": {
            instructor: "Saad Irfan khan",
            code: "AIC 3603",
        },
        "Technical and Business Writing": {
            instructor: "Aneela Kanwal",
            code: "CSC 1205",
        },
        "Computer Networks and Data Communications": {
            instructor: "Rana Faisal Hayat",
            code: "CSC 3205",
        },
        "Artificial Intelligence": {
            instructor: "Muhammad Qasim",
            code: "CSC 4101",
        },
        "Web Technologies-I": {
            instructor: "Zubair Ahmed Chatta ",
            code: "CSC 4717",
        },
        "Android Application Development": {
            instructor: "Muhammad Azhar",
            code: "CSC 4802",
        },
        "Lab: Computer Networks and Data Communications": {
            instructor: "Shafaq Rasheed",
            code: "CSCL 3205",
        },
        "Lab: Artificial Intelligence": {
            instructor: "Muhammad Usama",
            code: "CSCL 4101",
        },
    },
};
