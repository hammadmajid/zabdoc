export type CourseInfo = {
    instructor: string;
    code: string;
};

export type ClassData = Record<string, CourseInfo>;

export type DataStructure = Record<string, ClassData>;

export const data: DataStructure = {
    "BsCS-1A": {
        "Calculus and Analytical Geometry": {
            "instructor": "Dr. Tazeen Athar",
            "code": "CSC 1101"
        },
        "English Composition and Comprehension": {
            "instructor": "Sana Jaffery",
            "code": "CSC 1102"
        },
        "Fundamentals of Programming": {
            "instructor": "Sobia Rasheed Malik",
            "code": "CSC 1103"
        },
        "Applied Physics": {
            "instructor": "Dr. Shahzad Latif",
            "code": "CSC 1107"
        },
        "Introduction to Computer Science": {
            "instructor": "Tooba Abdul Qudoos Khan",
            "code": "CSC 1108"
        },
        "Lab: Fundamentals of Programming": {
            "instructor": "Ammad Noor",
            "code": "CSCL 1103"
        },
        "Lab: Applied Physics": {
            "instructor": "Muhammad Irfan",
            "code": "CSCL 1107"
        },
        "Lab: Introduction to Computer Science": {
            "instructor": "Usama Irshad",
            "code": "CSCL 1108"
        }
    },
    "BsCS-1B": {
        "Calculus and Analytical Geometry": {
            "instructor": "Dr. Aqsa Zafar Abbasi",
            "code": "CSC 1101"
        },
        "English Composition and Comprehension": {
            "instructor": "Irrum Arif",
            "code": "CSC 1102"
        },
        "Fundamentals of Programming": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 1103"
        },
        "Applied Physics": {
            "instructor": "Dr. Shahzad Latif",
            "code": "CSC 1107"
        },
        "Introduction to Computer Science": {
            "instructor": "Nazish Bashir",
            "code": "CSC 1108"
        },
        "Lab: Fundamentals of Programming": {
            "instructor": "Ammad Noor",
            "code": "CSCL 1103"
        },
        "Lab: Applied Physics": {
            "instructor": "Muhammad Irfan",
            "code": "CSCL 1107"
        },
        "Lab: Introduction to Computer Science": {
            "instructor": "Muhammad Farooq",
            "code": "CSCL 1108"
        }
    },
    "BsCS-1C": {
        "Calculus and Analytical Geometry": {
            "instructor": "Dr. Aqsa Zafar Abbasi",
            "code": "CSC 1101"
        },
        "English Composition and Comprehension": {
            "instructor": "Irrum Arif",
            "code": "CSC 1102"
        },
        "Fundamentals of Programming": {
            "instructor": "Sobia Rasheed Malik",
            "code": "CSC 1103"
        },
        "Applied Physics": {
            "instructor": "Mamoon Riaz",
            "code": "CSC 1107"
        },
        "Introduction to Computer Science": {
            "instructor": "Nazish Bashir",
            "code": "CSC 1108"
        },
        "Lab: Fundamentals of Programming": {
            "instructor": "Usama Irshad",
            "code": "CSCL 1103"
        },
        "Lab: Applied Physics": {
            "instructor": "Shehwar Tanveer",
            "code": "CSCL 1107"
        },
        "Lab: Introduction to Computer Science": {
            "instructor": "Ammad Noor",
            "code": "CSCL 1108"
        }
    },
    "BsCS-2A": {
        "Pakistan Studies": {
            "instructor": "Naeem Ahmad",
            "code": "CSC 1109"
        },
        "Probability and Statistics": {
            "instructor": "Saira . Zahid",
            "code": "CSC 1206"
        },
        "Digital Logic Design": {
            "instructor": "Dr. Shahzad Latif",
            "code": "CSC 1207"
        },
        "Object Oriented Programming Techniques": {
            "instructor": "Aniqa",
            "code": "CSC 1208"
        },
        "Islamic Studies / Humanities": {
            "instructor": "Dr. Muhammad Hammad",
            "code": "CSC 1209"
        },
        "Communication and Presentation Skills": {
            "instructor": "Sarah Saba",
            "code": "CSC 2101"
        },
        "Understanding of Holy Quran-I": {
            "instructor": "Dr. Muhammad Hammad",
            "code": "CSC 3115"
        },
        "Lab: Digital Logic Design": {
            "instructor": "Muhammad Irfan",
            "code": "CSCL 1207"
        },
        "Lab: Object Oriented Programming Techniques": {
            "instructor": "Usama Irshad",
            "code": "CSCL 1208"
        }
    },
    "BsCS-2B": {
        "Pakistan Studies": {
            "instructor": "Saima Sabir",
            "code": "CSC 1109"
        },
        "Probability and Statistics": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 1206"
        },
        "Digital Logic Design": {
            "instructor": "Tariq Salim",
            "code": "CSC 1207"
        },
        "Object Oriented Programming Techniques": {
            "instructor": "Tehreem Saboor",
            "code": "CSC 1208"
        },
        "Islamic Studies / Humanities": {
            "instructor": "Dr. Muhammad Hammad",
            "code": "CSC 1209"
        },
        "Communication and Presentation Skills": {
            "instructor": "Muhammad Adeel Ashraf",
            "code": "CSC 2101"
        },
        "Understanding of Holy Quran-I": {
            "instructor": "Dr. Muhammad Hammad",
            "code": "CSC 3115"
        },
        "Lab: Digital Logic Design": {
            "instructor": "Shehwar Tanveer",
            "code": "CSCL 1207"
        },
        "Lab: Object Oriented Programming Techniques": {
            "instructor": "Maria Bibi",
            "code": "CSCL 1208"
        }
    },
    "BsCS-2C": {
        "Pakistan Studies": {
            "instructor": "Saima Sabir",
            "code": "CSC 1109"
        },
        "Probability and Statistics": {
            "instructor": "Dr. Tanveer Kifayat",
            "code": "CSC 1206"
        },
        "Digital Logic Design": {
            "instructor": "Tariq Salim",
            "code": "CSC 1207"
        },
        "Object Oriented Programming Techniques": {
            "instructor": "Tehreem Saboor",
            "code": "CSC 1208"
        },
        "Islamic Studies / Humanities": {
            "instructor": "Komal Musheer Siddiqui",
            "code": "CSC 1209"
        },
        "Communication and Presentation Skills": {
            "instructor": "Muhammad Adeel Ashraf",
            "code": "CSC 2101"
        },
        "Understanding of Holy Quran-I": {
            "instructor": "Komal Musheer Siddiqui",
            "code": "CSC 3115"
        },
        "Lab: Digital Logic Design": {
            "instructor": "Shehwar Tanveer",
            "code": "CSCL 1207"
        },
        "Lab: Object Oriented Programming Techniques": {
            "instructor": "Bushra Shafi",
            "code": "CSCL 1208"
        }
    },
    "BsCS-2D": {
        "Pakistan Studies": {
            "instructor": "Naeem Ahmad",
            "code": "CSC 1109"
        },
        "Probability and Statistics": {
            "instructor": "Dr. Tanveer Kifayat",
            "code": "CSC 1206"
        },
        "Digital Logic Design": {
            "instructor": "Mamoon Riaz",
            "code": "CSC 1207"
        },
        "Object Oriented Programming Techniques": {
            "instructor": "Ayyaz Altaf Mirza",
            "code": "CSC 1208"
        },
        "Islamic Studies / Humanities": {
            "instructor": "Komal Musheer Siddiqui",
            "code": "CSC 1209"
        },
        "Communication and Presentation Skills": {
            "instructor": "Sarah Saba",
            "code": "CSC 2101"
        },
        "Understanding of Holy Quran-I": {
            "instructor": "Komal Musheer Siddiqui",
            "code": "CSC 3115"
        },
        "Lab: Digital Logic Design": {
            "instructor": "Muhammad Farooq",
            "code": "CSCL 1207"
        },
        "Lab: Object Oriented Programming Techniques": {
            "instructor": "Maria Bibi",
            "code": "CSCL 1208"
        }
    },
    "BsCS-2E": {
        "Pakistan Studies": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 1109"
        },
        "Probability and Statistics": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 1206"
        },
        "Digital Logic Design": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 1207"
        },
        "Object Oriented Programming Techniques": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 1208"
        },
        "Islamic Studies / Humanities": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 1209"
        },
        "Communication and Presentation Skills": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 2101"
        },
        "Understanding of Holy Quran-I": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 3115"
        },
        "Lab: Digital Logic Design": {
            "instructor": "Dr. Saad Malik",
            "code": "CSCL 1207"
        },
        "Lab: Object Oriented Programming Techniques": {
            "instructor": "Dr. Saad Malik",
            "code": "CSCL 1208"
        }
    },
    "BsCS-3A": {
        "Discrete Mathematical Structures": {
            "instructor": "Dr. Tazeen Athar",
            "code": "CSC 1201"
        },
        "Multivariate Calculus": {
            "instructor": "Tayyab Imran",
            "code": "CSC 1202"
        },
        "Teachings of Holy Quran": {
            "instructor": "Raheela Yasmeen",
            "code": "CSC 1215"
        },
        "Data Structures and Algorithms": {
            "instructor": "Saira Shaheen",
            "code": "CSC 2102"
        },
        "Computer Organization and Assembly Language": {
            "instructor": "Muhammad Usman Ali",
            "code": "CSC 3105"
        },
        "HCI and Computer Graphics": {
            "instructor": "Sobia Rasheed Malik",
            "code": "CSC 3106"
        },
        "Management Principles": {
            "instructor": "Beenish Ambereen",
            "code": "CSC 4906"
        },
        "Lab: Data Structures and Algorithms": {
            "instructor": "Maria Bibi",
            "code": "CSCL 2102"
        },
        "Lab: Computer Organization and Assembly Language": {
            "instructor": "Sohail Ahmad",
            "code": "CSCL 3105"
        },
        "Lab: HCI and Computer Graphics": {
            "instructor": "Imran Ahmad Qureshi",
            "code": "CSCL 3106"
        }
    },
    "BsCS-3B": {
        "Discrete Mathematical Structures": {
            "instructor": "Dr. Tazeen Athar",
            "code": "CSC 1201"
        },
        "Multivariate Calculus": {
            "instructor": "Dr. Aqeel Ahmed",
            "code": "CSC 1202"
        },
        "Teachings of Holy Quran": {
            "instructor": "Raheela Yasmeen",
            "code": "CSC 1215"
        },
        "Data Structures and Algorithms": {
            "instructor": "Fakhar Ul Islam",
            "code": "CSC 2102"
        },
        "Computer Organization and Assembly Language": {
            "instructor": "Hasnain Tahir",
            "code": "CSC 3105"
        },
        "HCI and Computer Graphics": {
            "instructor": "Tooba Abdul Qudoos Khan",
            "code": "CSC 3106"
        },
        "Management Principles": {
            "instructor": "Sadia Mir Ahmed",
            "code": "CSC 4906"
        },
        "Lab: Data Structures and Algorithms": {
            "instructor": "Nazish Bashir",
            "code": "CSCL 2102"
        },
        "Lab: Computer Organization and Assembly Language": {
            "instructor": "Muhammad Ishfaq",
            "code": "CSCL 3105"
        },
        "Lab: HCI and Computer Graphics": {
            "instructor": "Sohail Ahmad",
            "code": "CSCL 3106"
        }
    },
    "BsCS-3C": {
        "Discrete Mathematical Structures": {
            "instructor": "Tayyab Imran",
            "code": "CSC 1201"
        },
        "Multivariate Calculus": {
            "instructor": "Kamran Suhaib",
            "code": "CSC 1202"
        },
        "Teachings of Holy Quran": {
            "instructor": "Raheela Yasmeen",
            "code": "CSC 1215"
        },
        "Data Structures and Algorithms": {
            "instructor": "Ayyaz Altaf Mirza",
            "code": "CSC 2102"
        },
        "Computer Organization and Assembly Language": {
            "instructor": "Ghaffar Ahmed",
            "code": "CSC 3105"
        },
        "HCI and Computer Graphics": {
            "instructor": "Muhammad Usman Ali",
            "code": "CSC 3106"
        },
        "Management Principles": {
            "instructor": "Arslan Khurshid",
            "code": "CSC 4906"
        },
        "Lab: Data Structures and Algorithms": {
            "instructor": "Adeel Ahmed",
            "code": "CSCL 2102"
        },
        "Lab: Computer Organization and Assembly Language": {
            "instructor": "Muhammad Ishfaq",
            "code": "CSCL 3105"
        },
        "Lab: HCI and Computer Graphics": {
            "instructor": "Sohail Ahmad",
            "code": "CSCL 3106"
        }
    },
    "BsCS-4A": {
        "Pakistan Studies": {
            "instructor": "Adeela Iftikhar",
            "code": "CSC 1109"
        },
        "Teachings of Holy Quran": {
            "instructor": "Raheela Yasmeen",
            "code": "CSC 1215"
        },
        "Database Systems": {
            "instructor": "Dr. Saad Malik",
            "code": "CSC 2203"
        },
        "Finite Automata Theory and Formal Languages": {
            "instructor": "Muhammad Nadeem Khokhar",
            "code": "CSC 2204"
        },
        "Linear Algebra": {
            "instructor": "ABDUL MOQEET",
            "code": "CSC 2206"
        },
        "Computer Architecture": {
            "instructor": "Dr. Fadia Shah",
            "code": "CSC 3101"
        },
        "Design and Analysis of Algorithms": {
            "instructor": "Dr. Ghawar Said",
            "code": "CSC 3202"
        },
        "Lab: Database Systems": {
            "instructor": "Ammad Noor",
            "code": "CSCL 2203"
        }
    },
    "BsCS-4B": {
        "Pakistan Studies": {
            "instructor": "Adeela Iftikhar",
            "code": "CSC 1109"
        },
        "Teachings of Holy Quran": {
            "instructor": "Raheela Yasmeen",
            "code": "CSC 1215"
        },
        "Database Systems": {
            "instructor": "Ghaffar Ahmed",
            "code": "CSC 2203"
        },
        "Finite Automata Theory and Formal Languages": {
            "instructor": "Muhammad Nadeem Khokhar",
            "code": "CSC 2204"
        },
        "Linear Algebra": {
            "instructor": "Nasir shehzad",
            "code": "CSC 2206"
        },
        "Computer Architecture": {
            "instructor": "Dr. Fadia Shah",
            "code": "CSC 3101"
        },
        "Design and Analysis of Algorithms": {
            "instructor": "Hassan Ayaz",
            "code": "CSC 3202"
        },
        "Lab: Database Systems": {
            "instructor": "Ammad Noor",
            "code": "CSCL 2203"
        }
    },
    "BsCS-4C": {
        "Pakistan Studies": {
            "instructor": "Beenish Ambereen",
            "code": "CSC 1109"
        },
        "Teachings of Holy Quran": {
            "instructor": "Raheela Yasmeen",
            "code": "CSC 1215"
        },
        "Database Systems": {
            "instructor": "Muhammad Qasim",
            "code": "CSC 2203"
        },
        "Finite Automata Theory and Formal Languages": {
            "instructor": "Muhammad Taseer ul Islam",
            "code": "CSC 2204"
        },
        "Linear Algebra": {
            "instructor": "Ayesha Sahreen",
            "code": "CSC 2206"
        },
        "Computer Architecture": {
            "instructor": "Awais Mahmood",
            "code": "CSC 3101"
        },
        "Design and Analysis of Algorithms": {
            "instructor": "Imran Ahmad Qureshi",
            "code": "CSC 3202"
        },
        "Lab: Database Systems": {
            "instructor": "Muhammad Ishfaq",
            "code": "CSCL 2203"
        }
    },
    "BsCS-4D": {
        "Pakistan Studies": {
            "instructor": "Adeela Iftikhar",
            "code": "CSC 1109"
        },
        "Teachings of Holy Quran": {
            "instructor": "Muhammad Hassaan Raza",
            "code": "CSC 1215"
        },
        "Database Systems": {
            "instructor": "Aniqa",
            "code": "CSC 2203"
        },
        "Finite Automata Theory and Formal Languages": {
            "instructor": "Dr. Ghawar Said",
            "code": "CSC 2204"
        },
        "Linear Algebra": {
            "instructor": "Ayesha Sahreen",
            "code": "CSC 2206"
        },
        "Computer Architecture": {
            "instructor": "Awais Mahmood",
            "code": "CSC 3101"
        },
        "Design and Analysis of Algorithms": {
            "instructor": "Dr. Shahzad Latif",
            "code": "CSC 3202"
        },
        "Lab: Database Systems": {
            "instructor": "Maria Bibi",
            "code": "CSCL 2203"
        }
    },
    "BsCS-4E": {
        "Pakistan Studies": {
            "instructor": "Beenish Ambereen",
            "code": "CSC 1109"
        },
        "Teachings of Holy Quran": {
            "instructor": "Muhammad Hassaan Raza",
            "code": "CSC 1215"
        },
        "Database Systems": {
            "instructor": "Aniqa",
            "code": "CSC 2203"
        },
        "Finite Automata Theory and Formal Languages": {
            "instructor": "Arfa Asaf",
            "code": "CSC 2204"
        },
        "Linear Algebra": {
            "instructor": "Dr. Sajjad Ahmed Ghauri",
            "code": "CSC 2206"
        },
        "Computer Architecture": {
            "instructor": "Muhammad Ibne Amin",
            "code": "CSC 3101"
        },
        "Design and Analysis of Algorithms": {
            "instructor": "Rana Faisal Hayat",
            "code": "CSC 3202"
        },
        "Lab: Database Systems": {
            "instructor": "Bushra Shafi",
            "code": "CSCL 2203"
        }
    },
    "BsCS-5A": {
        "Teachings of Holy Quran": {
            "instructor": "Muhammad Hassaan Raza",
            "code": "CSC 1215"
        },
        "Differential Equations": {
            "instructor": "Dr. Sajjad Ahmed Ghauri",
            "code": "CSC 2122"
        },
        "Graph Theory": {
            "instructor": "Dr. Tazeen Athar",
            "code": "CSC 2123"
        },
        "Operating Systems": {
            "instructor": "Muhammad Qasim",
            "code": "CSC 2205"
        },
        "Software Engineering": {
            "instructor": "Sagher Abbas",
            "code": "CSC 3109"
        },
        "Compiler Construction": {
            "instructor": "Abbas Amir",
            "code": "CSC 3201"
        },
        "Lab: Operating Systems": {
            "instructor": "Adeel Ahmed",
            "code": "CSCL 2205"
        }
    },
    "BsCS-5B": {
        "Teachings of Holy Quran": {
            "instructor": "Muhammad Hassaan Raza",
            "code": "CSC 1215"
        },
        "Differential Equations": {
            "instructor": "Nasir shehzad",
            "code": "CSC 2122"
        },
        "Graph Theory": {
            "instructor": "Muhammad Farooq (CS)",
            "code": "CSC 2123"
        },
        "Operating Systems": {
            "instructor": "Sagher Abbas",
            "code": "CSC 2205"
        },
        "Software Engineering": {
            "instructor": "Awais Mahmood",
            "code": "CSC 3109"
        },
        "Compiler Construction": {
            "instructor": "Muhammad Nadeem Khokhar",
            "code": "CSC 3201"
        },
        "Lab: Operating Systems": {
            "instructor": "Adeel Ahmed",
            "code": "CSCL 2205"
        }
    },
    "BsCS-5C": {
        "Teachings of Holy Quran": {
            "instructor": "Muhammad Hassaan Raza",
            "code": "CSC 1215"
        },
        "Differential Equations": {
            "instructor": "ABDUL MOQEET",
            "code": "CSC 2122"
        },
        "Graph Theory": {
            "instructor": "Kamran Suhaib",
            "code": "CSC 2123"
        },
        "Operating Systems": {
            "instructor": "Fakhar Ul Islam",
            "code": "CSC 2205"
        },
        "Software Engineering": {
            "instructor": "Tehreem Saboor",
            "code": "CSC 3109"
        },
        "Compiler Construction": {
            "instructor": "Hassan Ayaz",
            "code": "CSC 3201"
        },
        "Lab: Operating Systems": {
            "instructor": "Sohail Ahmad",
            "code": "CSCL 2205"
        }
    },
    "BsCS-6A": {
        "Technical and Business Writing": {
            "instructor": "Aneela Kanwal",
            "code": "CSC 1205"
        },
        "Teachings of Holy Quran": {
            "instructor": "DrHafiz Javed Ahmad",
            "code": "CSC 1215"
        },
        "Computer Networks and Data Communications": {
            "instructor": "Rana Faisal Hayat",
            "code": "CSC 3205"
        },
        "Artificial Intelligence": {
            "instructor": "Asif Mehmood",
            "code": "CSC 4101"
        },
        "Web Technologies-I": {
            "instructor": "Zubair Ahmed Chatta",
            "code": "CSC 4717"
        },
        "Android Application Development": {
            "instructor": "Muhammad Saddheer",
            "code": "CSC 4802"
        },
        "Lab: Computer Networks and Data Communications": {
            "instructor": "Adeel Ahmed",
            "code": "CSCL 3205"
        },
        "Lab: Artificial Intelligence": {
            "instructor": "Muhammad Ishfaq",
            "code": "CSCL 4101"
        }
    },
    "BsCS-6B": {
        "Technical and Business Writing": {
            "instructor": "Aneela Kanwal",
            "code": "CSC 1205"
        },
        "Teachings of Holy Quran": {
            "instructor": "DrHafiz Javed Ahmad",
            "code": "CSC 1215"
        },
        "Computer Networks and Data Communications": {
            "instructor": "Muhammad Taseer ul Islam",
            "code": "CSC 3205"
        },
        "Artificial Intelligence": {
            "instructor": "Dr. Sajid Ali Khan",
            "code": "CSC 4101"
        },
        "Web Technologies-I": {
            "instructor": "Zubair Ahmed Chatta",
            "code": "CSC 4717"
        },
        "Android Application Development": {
            "instructor": "Muhammad Saddheer",
            "code": "CSC 4802"
        },
        "Lab: Computer Networks and Data Communications": {
            "instructor": "Muhammad Farooq",
            "code": "CSCL 3205"
        },
        "Lab: Artificial Intelligence": {
            "instructor": "Maria Bibi",
            "code": "CSCL 4101"
        }
    },
    "BsCS-6C": {
        "Technical and Business Writing": {
            "instructor": "Sana Jaffery",
            "code": "CSC 1205"
        },
        "Teachings of Holy Quran": {
            "instructor": "DrHafiz Javed Ahmad",
            "code": "CSC 1215"
        },
        "Computer Networks and Data Communications": {
            "instructor": "Fakhar Ul Islam",
            "code": "CSC 3205"
        },
        "Artificial Intelligence": {
            "instructor": "Awais - Nawaz",
            "code": "CSC 4101"
        },
        "Web Technologies-I": {
            "instructor": "Rana Faisal Hayat",
            "code": "CSC 4717"
        },
        "Android Application Development": {
            "instructor": "Zubair Ahmed Chatta",
            "code": "CSC 4802"
        },
        "Lab: Computer Networks and Data Communications": {
            "instructor": "Mutahra Khalid Butt",
            "code": "CSCL 3205"
        },
        "Lab: Artificial Intelligence": {
            "instructor": "Mutahra Khalid Butt",
            "code": "CSCL 4101"
        }
    },
    "BsCS-7A": {
        "Professional Practices": {
            "instructor": "Tooba Abdul Qudoos Khan",
            "code": "CSC 4102"
        },
        "Final Year Project-I": {
            "instructor": "",
            "code": "CSC 4105"
        },
        "Parallel and Distributed Computing": {
            "instructor": "Muhammad Ibne Amin",
            "code": "CSC 4106"
        },
        "Foreign Languages": {
            "instructor": "Somia Ashraf",
            "code": "CSC 4601"
        },
        "Applied Data Mining": {
            "instructor": "Shakeel Ahmad",
            "code": "CSC 4703"
        }
    },
    "BsCS-7B": {
        "Professional Practices": {
            "instructor": "Hasnain Tahir",
            "code": "CSC 4102"
        },
        "Final Year Project-I": {
            "instructor": "",
            "code": "CSC 4105"
        },
        "Parallel and Distributed Computing": {
            "instructor": "Rana Faisal Hayat",
            "code": "CSC 4106"
        },
        "Foreign Languages": {
            "instructor": "Somia Ashraf",
            "code": "CSC 4601"
        },
        "Applied Data Mining": {
            "instructor": "Hashim Safdar",
            "code": "CSC 4703"
        }
    },
    "BsCS-7C": {
        "Professional Practices": {
            "instructor": "Muhammad Islam",
            "code": "CSC 4102"
        },
        "Final Year Project-I": {
            "instructor": "",
            "code": "CSC 4105"
        },
        "Parallel and Distributed Computing": {
            "instructor": "Asma Niaz Awan",
            "code": "CSC 4106"
        },
        "Foreign Languages": {
            "instructor": "Maleeha Saqib",
            "code": "CSC 4601"
        },
        "Applied Data Mining": {
            "instructor": "Asif Mehmood",
            "code": "CSC 4703"
        }
    },
    "BsCS-7D": {
        "Professional Practices": {
            "instructor": "Awais - Nawaz",
            "code": "CSC 4102"
        },
        "Final Year Project-I": {
            "instructor": "",
            "code": "CSC 4105"
        },
        "Parallel and Distributed Computing": {
            "instructor": "Dr. Danish Mahmood",
            "code": "CSC 4106"
        },
        "Foreign Languages": {
            "instructor": "Maleeha Saqib",
            "code": "CSC 4601"
        },
        "Applied Data Mining": {
            "instructor": "Muhammad Usman",
            "code": "CSC 4703"
        }
    },
    "BsCS-8A": {
        "Information Security": {
            "instructor": "Awais Mahmood",
            "code": "CSC 4201"
        },
        "Final Year Project-II": {
            "instructor": "",
            "code": "CSC 4205"
        },
        "Business and Technology Ethics": {
            "instructor": "Muhammad Islam",
            "code": "CSC 4501"
        },
        "Digital Image Processing": {
            "instructor": "Nazish Bashir",
            "code": "CSC 4706"
        },
        "Introduction to Blockchain Technology": {
            "instructor": "Safi Ullah",
            "code": "CSC 4722"
        }
    },
    "BsCS-8B": {
        "Information Security": {
            "instructor": "Dr. Danish Mahmood",
            "code": "CSC 4201"
        },
        "Final Year Project-II": {
            "instructor": "",
            "code": "CSC 4205"
        },
        "Business and Technology Ethics": {
            "instructor": "Saira Shaheen",
            "code": "CSC 4501"
        },
        "Digital Image Processing": {
            "instructor": "Hashim Safdar",
            "code": "CSC 4706"
        },
        "Introduction to Blockchain Technology": {
            "instructor": "Asma Niaz Awan",
            "code": "CSC 4722"
        }
    },
    "BsCS-8C": {
        "Information Security": {
            "instructor": "Muhammad Taseer ul Islam",
            "code": "CSC 4201"
        },
        "Final Year Project-II": {
            "instructor": "",
            "code": "CSC 4205"
        },
        "Business and Technology Ethics": {
            "instructor": "Saira Shaheen",
            "code": "CSC 4501"
        },
        "Digital Image Processing": {
            "instructor": "Asif Mehmood",
            "code": "CSC 4706"
        },
        "Introduction to Blockchain Technology": {
            "instructor": "Shahab Khan",
            "code": "CSC 4722"
        }
    },
    "Open": {
        "Fundamentals of Programming": {
            "instructor": "Amjad Saleem Rana",
            "code": "CSC 1103"
        },
        "Applied Physics": {
            "instructor": "Resham Siddique",
            "code": "CSC 1107"
        },
        "Algebraic Foundations and Applications in Computing": {
            "instructor": "Asim - Shabir",
            "code": "CSC 1110"
        },
        "Differential and Integral Calculus in Computing": {
            "instructor": "Amara Bibi",
            "code": "CSC 1111"
        },
        "Object Oriented Programming Techniques": {
            "instructor": "Syed Mehdi Abbas Shah",
            "code": "CSC 1208"
        },
        "Digital Logic Design": {
            "instructor": "Dr. Shahzad Latif",
            "code": "CSC 2103"
        },
        "Computer Organization and Assembly Language": {
            "instructor": "Muhammad Usman Ali",
            "code": "CSC 2201"
        },
        "Management Principles": {
            "instructor": "Beenish Ambereen",
            "code": "CSC 4603"
        },
        "Lab: Fundamentals of Programming": {
            "instructor": "Amjad Saleem Rana",
            "code": "CSCL 1103"
        },
        "Lab: Applied Physics": {
            "instructor": "Resham Siddique",
            "code": "CSCL 1107"
        },
        "Lab: Object Oriented Programming Techniques": {
            "instructor": "Syed Mehdi Abbas Shah",
            "code": "CSCL 1208"
        },
        "Lab: Digital Logic Design": {
            "instructor": "Muhammad Irfan",
            "code": "CSCL 2103"
        },
        "Lab: Computer Organization and Assembly Language": {
            "instructor": "Muhammad Ishfaq",
            "code": "CSCL 2201"
        }
    }
};
