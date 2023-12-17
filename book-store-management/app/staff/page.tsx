import CreateStaffDialog from "@/components/staff/create-staff-dialog";
import { StaffTable } from "@/components/staff/staff-table";
import getAllStaff from "@/lib/staff/getAllStaff";
import { PagingProps, Staff } from "@/types";

const StaffManage = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const page = searchParams["page"] ?? "1";
  const active = searchParams["active"] ?? undefined;
  const search = searchParams["search"] ?? undefined;
  const staffsData: Promise<{ paging: PagingProps; data: Staff[] }> =
    getAllStaff({
      page: +page,
      isActive: active?.toString(),
      search: search?.toString(),
    });
  const staffs = await staffsData;
  const totalPage = Math.ceil(staffs.paging.total / staffs.paging.limit);
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách nhân viên</h1>
        <CreateStaffDialog />
      </div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <StaffTable data={staffs.data} totalPage={totalPage} />
      </div>
    </div>
  );
};

export default StaffManage;
