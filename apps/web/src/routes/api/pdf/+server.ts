// src/routes/api/pdf/+server.ts
import { json, type RequestHandler } from '@sveltejs/kit';
import { buildHTML } from '$lib/html';
import { documentSchema } from '@repo/types';
import { CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_API_TOKEN } from "$env/static/private";
import { Buffer } from "node:buffer";

export const POST: RequestHandler = async ({ request }) => {
  console.log('[PDF API] Incoming POST request');
  console.log('[PDF API] Request headers:', Object.fromEntries(request.headers.entries()));

  try {
    // Parse JSON from req
    console.log('[PDF API] Parsing JSON from request body');
    const jsonData = await request.json();
    console.log('[PDF API] JSON parsed successfully:', JSON.stringify(jsonData).substring(0, 200) + '...');

    // Validate with schema
    console.log('[PDF API] Validating data with documentSchema');
    const validated = documentSchema.parse(jsonData);
    console.log('[PDF API] Schema validation passed');
    console.log('[PDF API] Validated data keys:', Object.keys(validated));
    console.log('[PDF API] Number of students:', validated.Students?.length || 0);

    // Ensure required fields have defaults
    console.log('[PDF API] Adding default values to students');
    const validatedWithDefaults = {
      ...validated,
      Students: validated.Students.map(student => ({
        Name: student.Name || '',
        RegNo: student.RegNo || '',
      })),
    };
    console.log('[PDF API] Defaults applied. Student count:', validatedWithDefaults.Students.length);

    // Build HTML
    console.log('[PDF API] Building HTML from validated data');
    const html = buildHTML(validatedWithDefaults);
    console.log('[PDF API] HTML built successfully. Length:', html.length, 'characters');
    console.log('[PDF API] HTML preview:', html.substring(0, 150) + '...');

    // Send to Cloudflare
    console.log('[PDF API] Preparing Cloudflare API request');
    console.log('[PDF API] Account ID:', CLOUDFLARE_ACCOUNT_ID);
    console.log('[PDF API] API Token present:', !!CLOUDFLARE_API_TOKEN);

    const cloudflareUrl = `https://api.cloudflare.com/client/v4/accounts/${CLOUDFLARE_ACCOUNT_ID}/browser-rendering/pdf`;
    console.log('[PDF API] Cloudflare URL:', cloudflareUrl);
    console.log('[PDF API] Sending HTML to Cloudflare (size: ' + JSON.stringify({ html }).length + ' bytes)');

    const response = await fetch(cloudflareUrl, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${CLOUDFLARE_API_TOKEN}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ html })
    });

    console.log('[PDF API] Cloudflare response received');
    console.log('[PDF API] Response status:', response.status);
    console.log('[PDF API] Response headers:', Object.fromEntries(response.headers.entries()));

    if (!response.ok) {
      console.error('[PDF API] Cloudflare API returned non-ok status:', response.status);
      const errorData = await response.json();
      console.error('[PDF API] Cloudflare error details:', JSON.stringify(errorData));
      return json({ error: 'Cloudflare API failed', details: errorData }, { status: response.status });
    }

    console.log('[PDF API] Converting response to PDF buffer');
    const pdfBuffer = Buffer.from(await response.arrayBuffer());
    console.log('[PDF API] PDF buffer created. Size:', pdfBuffer.length, 'bytes');

    console.log('[PDF API] Returning PDF response');
    return new Response(pdfBuffer, {
      headers: {
        'Content-Type': 'application/pdf',
        'Content-Length': pdfBuffer.length.toString(),
        'Cache-Control': 'public, max-age=3600',
      },
    });
  } catch (error) {
    console.error('[PDF API] Error caught in try-catch block');
    console.error('[PDF API] Error type:', error instanceof Error ? error.name : typeof error);
    console.error('[PDF API] Error message:', error instanceof Error ? error.message : String(error));
    console.error('[PDF API] Full error:', error);

    // Return validation error if zod failed
    if (error instanceof Error && error.name === 'ZodError') {
      console.error('[PDF API] ZodError - Invalid schema validation');
      return json(
        { error: 'Invalid data', details: error.message },
        { status: 400 }
      );
    }

    console.error('[PDF API] Unexpected error - returning 500');
    return json({ error: 'Internal server error' }, { status: 500 });
  }
};