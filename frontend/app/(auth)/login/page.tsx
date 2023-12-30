"use server"
import styles from './login.module.css'
import { Button } from '@/components/ui/button';
import Image from 'next/image';
import LoginForm from '@/components/login/loginForm';
import { auth, signOut } from '@/lib/auth/auth';


const Login = async () => {
    return (
        <div className={styles.container}>
            <div className={styles.leftColumn}>
                <Image src='/login-background.png' alt='' fill className={styles.image} />
            </div>

            <div className={styles.rightColumn}>
                <div className={styles.rightHeader}>
                    <h1 className={styles.title}>Đăng nhập vào tài khoản</h1>
                    <h5 className={styles.description}>Xem những gì đang xảy ra với doanh nghiệp của bạn</h5>
                </div>

                <div className={styles.formContainer}>
                    <LoginForm />
                </div>
            </div>
        </div>
    );
};

export default Login;
