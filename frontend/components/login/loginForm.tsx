"use client";
import { useFormState } from "react-dom";
import { Input } from "../ui/input";
import { login } from "@/lib/auth/action";
import styles from './loginform.module.css'
import { Button } from "../ui/button";
import { Checkbox } from "../ui/checkbox";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useToast } from "../ui/use-toast";

const LoginForm = () => {
    const [state, formAction] = useFormState(login, undefined)
    const [isRemember, setIsRemember] = useState(false)
    const [toastTrigger, setToastTrigger] = useState(false)
    const { toast } = useToast()
    useEffect(() => {
        if (state?.error)
            toast({
                title: "Đăng nhập thất bại",
                description: "Vui lòng kiểm tra lại email và mật khẩu của bạn",
            });

    }, [state])

    return (
        <form onSubmit={() => setToastTrigger(!toastTrigger)} action={formAction}>
            <div className={styles.inputContainer}>
                <div className={styles.rightHeader} >
                    <div className='flex flex-col gap-1'>
                        <p>Email</p>
                        <Input
                            type='email'
                            placeholder='Email'
                            name="email"
                        />
                    </div>

                    <div className='flex flex-col gap-1'>
                        <p>Mật khẩu</p>
                        <Input
                            type='password'
                            placeholder='Mật khẩu'
                            name="password"
                        />
                    </div>
                </div>

                <div className={styles.rememberRow}>
                    <div className='flex flex-row items-center gap-3'>
                        <Checkbox checked={isRemember} onCheckedChange={(value: any) => setIsRemember(value)} />
                        <p>Ghi nhớ tôi</p>
                    </div>
                    <p className='text-[cornflowerblue]'>Quên mật khẩu?</p>
                </div>
            </div>
            <Button className={styles.button}>Đăng nhập</Button>
        </form>

    );
}

export default LoginForm;