"use client"
import { useParams, redirect } from "next/navigation"

export default function Page() {
    const params = useParams<{ id: string }>()
    redirect(process.env.NEXT_PUBLIC_REDIRECT_BASE_URL! + "/" + params.id)
}