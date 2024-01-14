"use client";
import EditDialog from "@/components/customer/edit";
import { InvoiceTable } from "@/components/customer/invoice-table";
import Loading from "@/components/loading";
import CustomerDetailSkeleton from "@/components/skeleton/customer-detail";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { endPoint } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import getCustomer from "@/lib/customer/getCustomer";
import { includesRoles } from "@/lib/utils";
import React from "react";
import { useSWRConfig } from "swr";

const CustomerDetail = ({ params }: { params: { customerId: string } }) => {
  const {
    data,
    isLoading,
    isError,
    mutate: mutateCustomer,
  } = getCustomer(params.customerId);
  const { currentUser } = useCurrentUser();
  const { mutate } = useSWRConfig();
  if (isError) return <div>Failed to load</div>;
  else if (isLoading) {
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
          <CustomerDetailSkeleton />
          <Card>
            <CardContent className="p-6 flex flex-col   gap-4">
              {/* TODO: invoice table */}
              <InvoiceTable customerId={params.customerId} />
            </CardContent>
          </Card>
        </div>
      </div>
    );
  } else
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
          <div className="flex justify-between">
            <h1 className="xl:text-3xl text-2xl">{data?.name}</h1>
            {!currentUser ||
            (currentUser &&
              includesRoles({
                currentUser: currentUser,
                allowedFeatures: ["CUSTOMER_UPDATE_INFO"],
              })) ? (
              <EditDialog
                customer={data}
                refresh={() => {
                  mutate(`${endPoint}/v1/customers/${data.id}`);
                }}
              />
            ) : null}
          </div>

          <Card>
            <CardContent className="p-6 flex flex-col   gap-4">
              <div className="flex gap-4  flex-col">
                <div className="flex gap-4 lg:flex-row flex-col">
                  <div className="basis-2/3">
                    <Label htmlFor="name">Mã khách hàng</Label>
                    <Input id="name" readOnly value={data?.id}></Input>
                  </div>
                  <div className="basis-1/3">
                    <Label htmlFor="phone">Điện thoại</Label>
                    <Input id="phone" readOnly value={data?.phone}></Input>
                  </div>
                </div>
                <div className="flex gap-4 lg:flex-row flex-col">
                  <div className="basis-2/3">
                    <Label htmlFor="email">Email</Label>
                    <Input id="email" readOnly value={data?.email}></Input>
                  </div>
                  <div className="basis-1/3">
                    <Label htmlFor="diem">Điểm tích luỹ</Label>
                    <Input id="diem" readOnly value={data?.point}></Input>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6 flex flex-col   gap-4">
              {/* TODO: invoice table */}
              <InvoiceTable customerId={params.customerId} />
            </CardContent>
          </Card>
        </div>
      </div>
    );
};

export default CustomerDetail;
