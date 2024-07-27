import ChangePasswordForm from "@/components/forgot-password/change_password_form";
import { Metadata } from "next";
import Image from "next/image";

export const metadata: Metadata = {
  title: "Đổi mật khẩu",
};
const ChangePassword = ({ params }: { params: { token: string } }) => {
  return (
    <div className="bg-auth-background bg-right bg-cover bg-no-repeat flex flex-1 flex-col p-8">
      <div className="flex flex-row align-middle">
        <Image
          src="/android-chrome-192x192.png"
          alt="logo"
          width={50}
          height={50}
        ></Image>
        <p className={`text-lg ml-2 font-semibold self-center`}>Book Store</p>
      </div>
      <div className="flex flex-1 flex-col justify-center"> 
        <div className="flex flex-row self-center gap-36">
          <Image
            src="/auth-image-2.png"
            alt="image-2"
            width={500}
            height={500}
            className="lg:block hidden"
          />
          <ChangePasswordForm token={params.token} />
        </div>
      </div>
    </div>
  );
};

export default ChangePassword;
