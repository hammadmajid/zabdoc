export type CourseInfo = {
    instructor: string;
    code: string;
    icon: string;
};

export type ClassData = Record<string, CourseInfo>;

export type DataStructure = {
    "BsCS-5C": ClassData;
};

export const data: DataStructure = {
    "BsCS-5C": {
        "Differential Equations": {
            instructor: "Abdul Moqeet",
            code: "CSC 2122",
            icon: "calculator",
        },
        "Graph Theory": {
            instructor: "Kamran Suhaib",
            code: "CSC 2123",
            icon: "git-branch",
        },
        "Operating Systems": {
            instructor: "Fakhar Ul Islam",
            code: "CSC 2205",
            icon: "hard-drive",
        },
        "Software Engineering": {
            instructor: "Tehreem Saboor",
            code: "CSC 3109",
            icon: "wrench",
        },
        "Compiler Construction": {
            instructor: "Hassan Ayaz",
            code: "CSC 3201",
            icon: "code",
        },
        "Lab: Operating Systems": {
            instructor: "Sohail Ahmad",
            code: "CSCL 2205",
            icon: "beaker",
        },
    },
};
