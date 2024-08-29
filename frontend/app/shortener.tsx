"use client"
import * as React from "react"

import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { CopyCheck, CopyIcon, Loader2 } from "lucide-react"
import { cn } from "@/lib/utils"
import Link from "next/link"

export function Shorten() {
    const baseUrl = process.env.NEXT_PUBLIC_REDIRECT_BASE_URL!
    const [url, setUrl] = React.useState("")
    const [shortUrl, setShortUrl] = React.useState("")
    const [copyClicked, setCopyClicked] = React.useState(false)
    const [validUrl, setValidUrl] = React.useState(true)
    const [loading, setLoading] = React.useState(false)
    const [isError, setIsError] = React.useState(false)

    async function handleShorten() {
        setShortUrl("")
        if (!isValidURL(url)) {
            setValidUrl(false)
            return
        }
        setLoading(true)
        try {
            const body = { "long_url": url }
            const response = await fetch(`${baseUrl}/shorten`, {
                body: JSON.stringify(body),
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                }
            });
            if (!response.ok) {
                setIsError(true)
                setLoading(false)
            }
            const json = await response.json();
            if (!json.short_id) {
                setIsError(true)
                setLoading(false)
            }

            setShortUrl(json.short_id)
            setLoading(false);
        } catch (error) {
            console.error(error)
            setIsError(true)
            setLoading(false)
        }
    }

    function handleCopyClicked() {
        setCopyClicked(true)
        setTimeout(() => setCopyClicked(false), 2000)
    }

    function isValidURL(urlString: string) {
        try {
            return Boolean(new URL(urlString));
        }
        catch (e) {
            return false;
        }
    }

    return (
        <Card className="w-[500px]">
            <CardHeader>
                <CardTitle>Paste the Long URL to be shortened</CardTitle>
            </CardHeader>
            <CardContent>
                <form>
                    <div className="grid w-full items-center gap-4 mb-4">
                        <div className="flex flex-col space-y-1.5">
                            <Input className={cn("appearance-none", !validUrl && "border-rose-300")} value={url} onChange={e => { setValidUrl(true); setUrl(e.target.value) }} id="name" placeholder="Enter the link here" />
                        </div>
                        {!validUrl && <p className="text-sm">Invalid url</p>}
                    </div>
                </form>
            </CardContent>
            <CardFooter className="flex justify-center">
                <Button onClick={handleShorten}>Shorten</Button>
            </CardFooter>
            <CardContent>
                <div className="flex justify-center">
                    {
                        loading && <Loader2 className="animate-spin"></Loader2>
                    }
                    {
                        shortUrl && (
                            <div className="flex items-center">
                                <Link target="_blank" href={`/${shortUrl}`}>
                                    <p className="underline text-md text-teal-600">{`${process.env.NEXT_PUBLIC_CLIENT_BASE_URL}/${shortUrl}`}</p>
                                </Link>
                                {
                                    !copyClicked &&
                                    <CopyIcon onClick={handleCopyClicked} className="ml-2 w-4 h-4 text-primary"></CopyIcon>
                                    ||
                                    <CopyCheck className="ml-2 w-4 h-4"></CopyCheck>
                                }
                            </div>
                        )
                    }
                </div>
            </CardContent>
        </Card>
    )
}
