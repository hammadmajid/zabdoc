<script lang="ts">
    import SEO from "$lib/components/seo.svelte";
</script>

<SEO
    title="about | zabdoc"
    description="Learn about zabdoc, an independent tool for SZABIST students to generate assignment and lab PDFs. Not affiliated with SZABIST."
    canonical="https://zabdoc.xyz/about"
    url="https://zabdoc.xyz/about"
/>

<div class="prose prose-neutral dark:prose-invert max-w-5xl mx-auto">
    <h1>About</h1>

    <p>
        <strong>zabdoc</strong> is a web application for
        <span class="border-b border-dotted border-foreground/50">SZABIST</span>
        students to generate assignment cover sheets and lab tasks in PDF format.
    </p>

    <p>
        Written in{" "}
        <a href="https://go.dev" target="_blank" rel="noreferrer">Go</a> and
        <a href="https://svelte.dev" target="_blank" rel="noreferrer">Svelte</a
        >, it provides a simple, responsive web form and produces
        institution-branded PDFs for assignment cover and lab tasks.
    </p>

    <h2 id="disclaimer">Disclaimer</h2>

    <p>
        <strong>Important:</strong> This project is{" "}
        <em>not sponsored, affiliated, endorsed, or approved by</em> Shaheed Zulfikar
        Ali Bhutto Institute of Science and Technology (SZABIST) Islamabad or any
        of its other campuses. This is an independent, unofficial tool created to
        assist students.
    </p>

    <p>
        <strong>No Warranty:</strong> This program is distributed in the hope
        that it will be useful, but <em>WITHOUT ANY WARRANTY</em>; without even
        the implied warranty of{" "}
        <em>MERCHANTABILITY</em> or <em>FITNESS FOR A PARTICULAR PURPOSE</em>.
        See the
        <a
            href="https://www.gnu.org/licenses/gpl-3.0.html"
            target="_blank"
            rel="noopener noreferrer"
        >
            GNU General Public License</a
        >
        for more details.
    </p>

    <p>
        The authors and contributors of this software cannot be held accountable
        for any issues, damages, or problems that may arise from the use of this
        application.
    </p>

    <h2>Terms of Use</h2>

    <ul>
        <li>
            This tool is provided for educational and convenience purposes only
        </li>
        <li>
            Users are responsible for ensuring the accuracy of information
            entered
        </li>
        <li>
            The generated PDFs should comply with your institution's submission
            requirements
        </li>
        <li>Use of this service is at your own risk and discretion</li>
    </ul>

    <h2>Privacy Information</h2>

    <p>
        We respect your privacy and are committed to protecting your personal
        information:
    </p>

    <ul>
        <li>
            <strong>Data Collection:</strong> We only collect the information you
            provide in the form to generate PDF document
        </li>
        <li>
            <strong>Data Usage:</strong> Your information is used solely to generate
            the requested PDF document
        </li>
        <li>
            <strong>Data Sharing:</strong> We do not share, sell, or distribute your
            personal information to third parties
        </li>
        <li>
            <strong>Cookies:</strong> We may use local storage to save your preferences
            for convenience
        </li>
    </ul>

    <h2>Support</h2>

    <p>
        For questions or support, please open an issue on{" "}
        <a
            href="https://github.com/hammadmajid/zabdoc/issues"
            target="_blank"
            rel="noopener noreferrer"
        >
            GitHub
        </a>.
    </p>
</div>
