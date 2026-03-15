<script lang="ts">
    import SEO from "$lib/components/seo.svelte";
</script>

<SEO
    title="Privacy Policy | zabdoc"
    description="Privacy Policy for zabdoc. Learn how we handle your data, credentials, and personal information."
    canonical="https://zabdoc.xyz/privacy"
    url="https://zabdoc.xyz/privacy"
/>

<div class="mx-auto space-y-8">
    <div class="prose prose-neutral dark:prose-invert prose-headings:font-black prose-headings:uppercase prose-headings:tracking-tight prose-a:text-primary prose-a:font-bold prose-strong:font-black max-w-none">
        <h1>Privacy Policy</h1>

        <p class="text-lg">
            zabdoc respects your privacy. This Privacy Policy explains how we collect, use, disclose, and safeguard
            information when you use our Service. Please read this policy carefully. If you do not agree with our
            policies and practices, please do not use our Service.
        </p>

        <h2>1. Information We Collect</h2>

        <h3>Information You Provide Directly</h3>

        <p>We collect information you voluntarily provide when using zabdoc, including:</p>
        <ul>
            <li><strong>Student Information:</strong> Name, registration number, class, and email</li>
            <li><strong>Course Details:</strong> Course name, instructor, course code, semester</li>
            <li><strong>Document Information:</strong> Assignment/lab number, marks, due dates, descriptions</li>
            <li><strong>Images:</strong> Lab task images and screenshots you upload</li>
            <li><strong>ZabDesk Credentials:</strong> Username and password (for scraping only, see below)</li>
        </ul>

        <h3>Information Automatically Collected</h3>

        <p>When you use zabdoc, we automatically collect:</p>
        <ul>
            <li>Your IP address and browser information</li>
            <li>The pages you visit and time spent on each page</li>
            <li>Error logs and debugging information</li>
            <li>Analytics data to improve the Service</li>
        </ul>

        <h2>2. Local Storage & Browser Data</h2>

        <p>
            <strong>Student Details Storage:</strong> Your student name and registration number are saved
            <strong>locally in your browser</strong> using browser storage (localStorage/IndexedDB).
            This data:
        </p>
        <ul>
            <li>Never leaves your device automatically</li>
            <li>Persists across browser sessions for your convenience</li>
            <li>Can be deleted anytime via the Settings page</li>
            <li>Is specific to your device and browser profile</li>
        </ul>

        <p>
            <strong>ZabDesk Credentials:</strong> Your ZabDesk username and password are <strong>NOT</strong> saved
            locally by zabdoc. They are transmitted securely to our backend only when you submit the scraping form.
        </p>

        <h2>3. ZabDesk Credentials - Special Handling</h2>

        <p>
            <strong>How we use ZabDesk credentials:</strong>
        </p>
        <ul>
            <li>You provide credentials only when scraping ZabDesk data</li>
            <li>Credentials are sent securely (HTTPS) to our backend API</li>
            <li>The backend uses credentials to authenticate with ZabDesk and fetch your data</li>
            <li>Credentials are <strong>immediately discarded after the request completes</strong></li>
            <li>We do not store, log, or retain your credentials</li>
        </ul>

        <p>
            <strong>We are not responsible for:</strong>
        </p>
        <ul>
            <li>How ZabDesk handles authentication requests</li>
            <li>Compromised ZabDesk accounts (use a strong password)</li>
            <li>Third-party credential harvesting or phishing attacks</li>
        </ul>

        <p>
            <strong>Security Recommendation:</strong> If you're concerned about credential security, you can:
        </p>
        <ul>
            <li>Change your ZabDesk password after using the scraper</li>
            <li>Use a temporary password if available through ZabDesk</li>
            <li>Contact SZABIST IT if you suspect unauthorized access</li>
        </ul>

        <h2>4. Scraped Data - Attendance & Marks</h2>

        <p>
            When you scrape ZabDesk data:
        </p>
        <ul>
            <li>Your attendance records and marks are fetched from ZabDesk servers</li>
            <li>Data is displayed in your browser for your review</li>
            <li>Data may be temporarily processed on our backend to fetch and format it</li>
            <li>Data is <strong>not permanently stored</strong> on our servers</li>
            <li>Data is cleared when you refresh the page or navigate away</li>
        </ul>

        <p>
            <strong>Data Retention:</strong> Scraped attendance and marks data is not retained after your session ends.
        </p>

        <h2>5. Data Sharing & Third Parties</h2>

        <p>
            <strong>We do NOT share your personal information with third parties</strong> except when required by law
            or court order. We do not:
        </p>
        <ul>
            <li>Sell your data to advertisers or marketers</li>
            <li>Share credentials with anyone</li>
            <li>Use your data for targeted advertising</li>
            <li>Provide your information to SZABIST (except as required by law)</li>
            <li>Partner with data brokers</li>
        </ul>

        <p>
            <strong>Essential Service Providers:</strong> We may share minimal information with:
        </p>
        <ul>
            <li>Cloud hosting providers (Heroku, Cloudflare) - infrastructure only</li>
            <li>Analytics providers - anonymized usage data only</li>
            <li>Law enforcement - only when compelled by legal process</li>
        </ul>

        <h2>6. Cookies & Tracking</h2>

        <p>
            zabdoc uses minimal cookies. We may use:
        </p>
        <ul>
            <li><strong>Session Cookies:</strong> For authentication and user sessions</li>
            <li><strong>Preference Cookies:</strong> To remember your settings</li>
            <li><strong>Analytics Cookies:</strong> To track usage patterns (anonymized)</li>
        </ul>

        <p>
            You can disable cookies in your browser settings, though this may affect functionality.
        </p>

        <h2>7. Logging & Monitoring</h2>

        <p>
            For debugging and operational purposes, we may log:
        </p>
        <ul>
            <li>HTTP request/response data and errors</li>
            <li>Error stacktraces and debugging information</li>
            <li>Performance metrics and system health data</li>
            <li>"Wide events" - structured logs that may contain user-provided fields</li>
        </ul>

        <p>
            <strong>Important:</strong> These logs may temporarily contain information you submitted (form data,
            course details) for troubleshooting purposes. Logs are:
        </p>
        <ul>
            <li>Retained for a limited period (typically 30 days)</li>
            <li>Not accessible to the public</li>
            <li>Used only for debugging and security</li>
            <li>Never shared with third parties for marketing</li>
        </ul>

        <h2>8. Data Security</h2>

        <p>
            We implement reasonable security measures including:
        </p>
        <ul>
            <li>HTTPS encryption for all data in transit</li>
            <li>Secure backend infrastructure (Heroku)</li>
            <li>Regular security updates and patches</li>
            <li>Limited data retention</li>
        </ul>

        <p>
            <strong>However:</strong> No security system is impenetrable. We cannot guarantee absolute protection
            against all attacks or breaches. Use zabdoc at your own risk.
        </p>

        <h2>9. Your Privacy Rights</h2>

        <p>
            You have the right to:
        </p>
        <ul>
            <li><strong>Access:</strong> Request what data we have about you</li>
            <li><strong>Delete:</strong> Request deletion of your local data via Settings page</li>
            <li><strong>Opt-Out:</strong> Disable analytics tracking (if available)</li>
            <li><strong>Correction:</strong> Request correction of inaccurate data</li>
        </ul>

        <p>
            To exercise these rights, <a href="mailto:hammadmajid@pm.me">contact us</a>.
        </p>

        <h2>10. Children's Privacy</h2>

        <p>
            zabdoc is intended for SZABIST students (typically 18+). We do not knowingly collect information from
            children under 13. If we discover such information, we will delete it immediately.
        </p>

        <h2>11. International Data</h2>

        <p>
            zabdoc is hosted on servers in the United States (Heroku) and uses Cloudflare's global CDN. By using
            the Service, you consent to your data being transferred to and processed in the United States and other
            countries. Data protection laws in these countries may differ from Pakistan.
        </p>

        <h2>12. Changes to This Policy</h2>

        <p>
            We may update this Privacy Policy at any time. Changes take effect immediately upon posting.
            Your continued use of zabdoc constitutes acceptance of the updated policy.
        </p>

        <h2>13. Contact Us</h2>

        <p>
            If you have questions about this Privacy Policy or our data practices, please contact:
        </p>
        <ul>
            <li>Email: <a href="mailto:hammadmajid@pm.me">hammadmajid@pm.me</a></li>
            <li>GitHub: <a href="https://github.com/hammadmajid/zabdoc" target="_blank" rel="noreferrer">GitHub Issues</a></li>
        </ul>

        <h2>14. Compliance</h2>

        <p>
            While zabdoc is not subject to GDPR (not operated in the EU), we attempt to follow privacy best practices
            similar to GDPR principles where feasible. We respect user privacy and data minimization principles.
        </p>

        <hr class="my-8" />

        <p class="text-sm text-muted-foreground">
            Last updated: March 2026.
        </p>

        <p class="text-sm text-muted-foreground">
            See also: <a href="/terms" class="text-primary font-bold hover:underline">Terms of Use</a> | <a href="/about" class="text-primary font-bold hover:underline">About zabdoc</a>
        </p>
    </div>
</div>
