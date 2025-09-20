<script lang="ts">
    import SEO from '$lib/components/seo.svelte';
</script>

<SEO title="help | zabdoc"
     description="Help and FAQ for zabdoc. Learn how to generate assignment, lab task, and lab project PDFs for SZABIST students."
     canonical="https://zabdoc.xyz/help"
     url="https://zabdoc.xyz/help"
/>


<div class="prose prose-neutral dark:prose-invert max-w-5xl mx-auto">
    <h1>Help</h1>

    <p>
        <strong>zabdoc</strong> helps SZABIST students generate properly formatted
        assignment and lab task documents in PDF format. This guide explains how
        to use each project type effectively.
    </p>

    <h2>How It Works</h2>

    <p>
        zabdoc uses an intelligent auto-form system that simplifies document
        creation by automatically filling course codes and instructor
        information based on your class and course selection.
    </p>

    <h3>Auto-Form Features</h3>
    <ul>
        <li>
            <strong>Class Selection:</strong> Choose your class (e.g., BsCS-4C) from
            the dropdown
        </li>
        <li>
            <strong>Course Selection:</strong> Select your course from the available
            options for your class
        </li>
        <li>
            <strong>Auto-Fill:</strong> Course code and instructor name are automatically
            populated
        </li>
        <li>
            <strong>Local Storage:</strong> Your name, registration number, and class
            are saved for future use
        </li>
    </ul>

    <h2>Available Document Types</h2>

    <h3>Assignment</h3>
    <p>
        Generate assignment documents that include all required student and
        course information along with assignment details.
    </p>

    <h4>How to Use:</h4>
    <ol>
        <li>Click on <strong>Assignment</strong> from the home page</li>
        <li>Enter your student information (name and registration number)</li>
        <li>Select your class from the dropdown menu</li>
        <li>Choose your course (instructor and course code will auto-fill)</li>
        <li>Specify assignment number, total marks, and due date</li>
        <li>Click <strong>Generate Document</strong> to download your PDF</li>
    </ol>

    <h3>Lab Task</h3>
    <p>
        Create formatted lab task documents for laboratory assignments and
        practical work.
    </p>

    <h4>How to Use:</h4>
    <ol>
        <li>Click on <strong>Lab Task</strong> from the home page</li>
        <li>Fill in your student information</li>
        <li>Select your class and course (details auto-fill)</li>
        <li>Specify lab task number, marks, and submission date</li>
        <li>Add content with supporting images in each section.</li>
        <li>Generate and download your lab task document</li>
    </ol>

    <h3>Lab Project</h3>
    <p>
        <!-- TODO: add more information -->
        <i>Work in Progress. More information will be added in the future.</i>
    </p>

    <h2>Missing Class or Course Information</h2>

    <p>
        If your class or course is not available in the dropdown menus, you have
        several options to request addition:
    </p>

    <h3>Option 1: Email Request</h3>
    <ul>
        <li>
            Contact <a href="mailto:hammadmajid@pm.me">hammadmajid@pm.me</a> with
            your request
        </li>
    </ul>

    <h3>Option 2: GitHub Issue</h3>
    <ul>
        <li>
            Open an issue on <a
                href="https://github.com/hammadmajid/zabdoc/issues"
                target="_blank"
                rel="noopener noreferrer">GitHub</a
            >
        </li>
        <li>Use the "Request Class/Course Addition" template</li>
    </ul>

    <h3>Option 3: Contribute Yourself</h3>
    <ul>
        <li>Add the information yourself by opening a pull request</li>
        <li>
            Follow the guide at <a
                href="https://github.com/hammadmajid/zabdoc/tree/main/docs/adding-course.md"
                target="_blank"
                rel="noopener noreferrer">Adding Course Documentation</a
            >
        </li>
    </ul>

    <h3>Required Information</h3>
    <p>
        <strong>Important:</strong> When requesting a new class, please provide
        information for <em>all courses</em> offered to that specific class.
    </p>

    <p>For each course, include:</p>
    <ul>
        <li>Class name (e.g., BsCS-8A)</li>
        <li>Course name (e.g., Software Engineering)</li>
        <li>Course code (e.g., CSC 3301)</li>
        <li>Instructor name</li>
    </ul>

    <p>New classes and courses are typically added within 24-48 hours.</p>

    <h2>Common Issues & Solutions</h2>

    <h3>PDF Not Downloading</h3>
    <ul>
        <li>Check that all required fields are filled correctly</li>
        <li>Ensure your browser allows downloads from this site</li>
        <li>Try refreshing the page and submitting again</li>
        <li>Disable browser extensions that might block downloads</li>
    </ul>

    <h3>Browser Compatibility & Requirements</h3>
    <ul>
        <li>
            <strong>Supported Browsers:</strong> Chrome 90+, Firefox 88+, Safari
            14+, Edge 90+
        </li>
        <li>
            <strong>JavaScript:</strong> Must be enabled for full functionality
        </li>
        <li>
            <strong>Internet Connection:</strong> Required for PDF generation
        </li>
        <li>
            <strong>Local Storage:</strong> Used to save your student information
            for convenience
        </li>
        <li><strong>Clear browser cache if experiencing issues</strong></li>
    </ul>

    <h2>Frequently Asked Questions</h2>

    <h3>Is this tool officially approved by SZABIST?</h3>
    <p>
        No, this is an independent project created to help students. It is not
        affiliated with or endorsed by SZABIST. See our <a href="/about"
            >About page</a
        > for full disclaimer.
    </p>

    <h3>Can I save my information for future use?</h3>
    <p>
        Your student name, registration number, and class selection are
        automatically saved locally in your browser for convenience.
    </p>

    <h3>How can I delete my saved information?</h3>
    <p>
        To remove your locally stored student information, go to the <a
            href="/settings">Settings</a
        >
        page and click on <strong>Delete Local Data</strong>. This will clear
        your name, registration number, and class selection from your browser.
    </p>

    <h3>What if my course format doesn't match the template?</h3>
    <p>
        The templates are designed to work with standard SZABIST course formats.
        If your course or class is not available in the dropdown menus, see the
        "Missing Class or Course Information" section above for multiple ways to
        request addition. If you encounter other issues, please check with your
        instructor about acceptable submission formats.
    </p>

    <h3>Can I modify the generated PDF?</h3>
    <p>
        The generated PDFs are in final format. If you need changes, you'll need
        to regenerate the document with corrected information.
    </p>

    <h2>Getting Support</h2>

    <p>If you encounter technical issues or have suggestions:</p>

    <ul>
        <li>Check this help page for common solutions</li>
        <li>
            Review the <a href="/about">About page</a> for additional information
        </li>
        <li>
            Report bugs or request features on <a
                href="https://github.com/hammadmajid/zabdoc/issues"
                target="_blank"
                rel="noopener noreferrer">GitHub Issues</a
            >
        </li>
        <li>
            For class/course additions, see the "Missing Class or Course
            Information" section above
        </li>
    </ul>

    <p>
        <strong>Note:</strong> This tool is provided as-is for educational purposes.
        Always verify that generated documents meet your instructor's requirements
        before submission.
    </p>
</div>
