import subprocess
import os
import signal
import atexit
import time

go_processes = []

go_files = {
    "api/gateway": "gateway.go",
    "rpc/user": "user.go",
    "rpc/question": "question.go",
    "rpc/judge": "judge.go",
}

compiler_flag = True


def build_go_file(go_src_path, go_src_file):
    exe_name = os.path.splitext(go_src_file)[0] + ".exe"
    if not compiler_flag:
        return exe_name
    print(f"编译 {go_src_file} -> {exe_name}")
    result = subprocess.run(
        [
            "go",
            "build",
            "-o",
            f"./{go_src_path}/{exe_name}",
            f"./{go_src_path}/{go_src_file}",
        ],
        capture_output=True,
        text=True,
    )
    if result.returncode != 0:
        print(f"编译失败：{result.stderr}")
        return None
    return exe_name


def start_exe(go_src_path, exe_name, cwd):
    return subprocess.Popen(
        [f"./{go_src_path}/{exe_name}"],
        creationflags=subprocess.CREATE_NEW_PROCESS_GROUP,
        cwd=cwd,
    )


def cleanup():
    print("清理中，关闭所有 Go 程序...")
    for proc in go_processes:
        try:
            proc.send_signal(signal.CTRL_BREAK_EVENT)
            proc.wait(timeout=5)
            print(f"成功关闭 {proc.pid}")
        except Exception as e:
            print(f"关闭进程失败：{e}")


atexit.register(cleanup)

if not os.path.exists("./output"):
    os.makedirs("./output")

for go_src_path, go_src_file in go_files.items():
    exe_name = build_go_file(go_src_path, go_src_file)
    if exe_name:
        proc = start_exe(go_src_path, exe_name, f"./{go_src_path}")
        go_processes.append(proc)

print("所有 Go 程序已启动。Python 程序继续运行...\n按 Ctrl+C 退出")

try:
    while True:
        time.sleep(9999999)
except KeyboardInterrupt:
    print("收到 Ctrl+C，正在退出...")
