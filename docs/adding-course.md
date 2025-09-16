# Adding Classes and Courses to zabdoc

This guide explains how to add new classes and courses to the zabdoc application by contributing directly to the codebase.

## Overview

Course and class data in zabdoc is stored in the file `worker/src/lib/data/data.ts`. This file contains a TypeScript object that maps class names to their courses, along with instructor and course code information.

## Understanding the Data Structure

The data structure consists of:

```typescript
export type CourseInfo = {
    instructor: string;
    code: string;
};

export type ClassData = Record<string, CourseInfo>;

export type DataStructure = {
    "BsCS-4C": ClassData;
    "BsCS-6A": ClassData;
    // Add new classes here
};
```

## Step-by-Step Guide

### 1. Fork the Repository

1. Go to [https://github.com/hammadmajid/zabdoc](https://github.com/hammadmajid/zabdoc)
2. Click the "Fork" button in the top-right corner
3. Clone your fork to your local machine:

   ```bash
   git clone https://github.com/<YOUR_USERNAME>/zabdoc.git
   cd zabdoc
   ```

### 2. Navigate to the Data File

Open the file `worker/src/lib/data/data.ts` in your preferred text editor.

### 3. Adding a New Class

#### Option A: Adding a Completely New Class

1. **Update the DataStructure type** (around line 8-11):

   ```typescript
   export type DataStructure = {
       "BsCS-4C": ClassData;
       "BsCS-6A": ClassData;
       "BsCS-8A": ClassData;  // Add your new class here
   };
   ```

2. **Add the class data** to the main data object (around line 13):

   ```typescript
   export const data: DataStructure = {
       "BsCS-4C": {
           // existing courses...
       },
       "BsCS-6A": {
           // existing courses...
       },
       "BsCS-8A": {  // Your new class
           "Software Engineering": {
               instructor: "Dr. John Smith",
               code: "CSC 3301",
           },
           "Database Management Systems": {
               instructor: "Prof. Jane Doe",
               code: "CSC 3201",
           },
           // Add all courses for this class
       },
   };
   ```

#### Option B: Adding Courses to an Existing Class

Simply add new course entries to the existing class object:

```typescript
"BsCS-4C": {
    // existing courses...
    "New Course Name": {
        instructor: "Instructor Name",
        code: "COURSE CODE",
    },
},
```

### 4. Course Entry Format

Each course entry should follow this exact format:

```typescript
"Course Name": {
    instructor: "Instructor Name",
    code: "COURSE CODE",
},
```

**Important Guidelines:**

- **Course Name**: Use the full, official course name as it appears in university documents
- **Instructor**: Include titles (Dr., Prof., Mr., Ms.) if commonly used
- **Code**: Use the exact course code format (e.g., "CSC 2203", "CSCL 2203")
- **Consistency**: Maintain consistent formatting with existing entries

### 5. Example: Complete Class Addition

Here's a complete example of adding a new class:

```typescript
export type DataStructure = {
    "BsCS-4C": ClassData;
    "BsCS-6A": ClassData;
    "BsCS-8A": ClassData;  // New class added to type
};

export const data: DataStructure = {
    "BsCS-4C": {
        // existing BsCS-4C courses...
    },
    "BsCS-6A": {
        // existing BsCS-6A courses...
    },
    "BsCS-8A": {  // New class implementation
        "Software Engineering": {
            instructor: "Dr. Ahmed Hassan",
            code: "CSC 3301",
        },
        "Database Management Systems": {
            instructor: "Prof. Sarah Khan",
            code: "CSC 3201",
        },
        "Computer Graphics": {
            instructor: "Mr. Ali Raza",
            code: "CSC 3401",
        },
        "Lab: Software Engineering": {
            instructor: "Ms. Fatima Sheikh",
            code: "CSCL 3301",
        },
        "Lab: Database Management Systems": {
            instructor: "Mr. Usman Ahmed",
            code: "CSCL 3201",
        },
    },
};
```

### 6. Testing Your Changes

1. **Install dependencies** (if you haven't already):

   ```bash
   cd worker
   npm install
   ```

2. **Start the development server**:

   ```bash
   npm run dev
   ```

3. **Test the changes**:
   - Navigate to the assignment or lab task page
   - Select your newly added class from the dropdown
   - Verify that all courses appear correctly
   - Ensure instructor and course code auto-fill properly

### 7. Commit and Submit

1. **Commit your changes**:

   ```bash
   git add worker/src/lib/data/data.ts
   git commit -m "Add BsCS-8A class with all courses"
   ```

2. **Push to your fork**:

   ```bash
   git push origin main
   ```

3. **Create a Pull Request**:
   - Go to your fork on GitHub
   - Click "New Pull Request"
   - Provide a clear title: "Add [ClassName] with all courses"
   - In the description, list all the courses you've added

## Best Practices

### Data Quality

- **Accuracy**: Double-check all course codes and instructor names
- **Completeness**: Include ALL courses for a class, not just some
- **Consistency**: Follow the same naming conventions as existing entries
- **Official Information**: Use only verified, official course information

### Course Naming

- Use the official course name from university catalogs
- For lab courses, prefix with "Lab: " (e.g., "Lab: Database Systems")
- Maintain consistent capitalization

### Instructor Names

- Include appropriate titles (Dr., Prof., Mr., Ms.) when commonly used
- Use the name as it appears in official course documents
- Be consistent with existing entries in the same class

### Code Style

- Maintain proper TypeScript formatting
- Use double quotes for strings
- Include trailing commas where appropriate
- Follow the existing indentation pattern (4 spaces)

## Getting Help

If you encounter issues while contributing:

1. **Check the Console**: Look for TypeScript or JavaScript errors in the browser console
2. **Verify Syntax**: Ensure proper JSON/TypeScript syntax (commas, quotes, brackets)
3. **Ask for Help**: Open an issue on GitHub or contact <hammadmajid@pm.me>

## Review Process

After submitting your pull request:

1. **Manual Review**: A maintainer will review your changes
2. **Feedback**: You may receive requests for modifications
3. **Merge**: Once approved, your changes will be merged and deployed

Your contribution helps make zabdoc more useful for SZABIST students. Thank you for contributing!
