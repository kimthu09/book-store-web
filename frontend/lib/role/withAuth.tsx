// withAuth.tsx
import { ReactNode } from "react";
import { getUser } from "../auth/action";
import NoRole from "@/components/no-role";
import { isAdmin } from "../utils";
interface WithAuthProps {
  children: ReactNode;
}
type ComponentType = (props: any) => JSX.Element;

export const withAuth = (
  WrappedComponent: ComponentType,
  allowedFeatures: string[],
  needAdmin?: boolean
) => {
  return async ({ ...props }: any) => {
    const currentUser = await getUser();
    const json = JSON.stringify(currentUser);
    const user = JSON.parse(json);

    const features = user.data.role.features.map((item: any) => item.featureId);
    if (needAdmin) {
      const isAdminRole = isAdmin({ currentUser: currentUser });
      if (isAdminRole) {
        return <WrappedComponent {...props} />;
      } else {
        return <NoRole />;
      }
    } else if (
      currentUser &&
      allowedFeatures.every((item) => features.includes(item))
    ) {
      return <WrappedComponent {...props} />;
    } else {
      // Redirect or show an error message
      return <NoRole />;
    }
  };
};
