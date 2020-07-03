# frozen_string_literal: true

if child_pid = fork
  puts "parent process. pid: #{Process.pid}, child pid: #{child_pid}"
  # => parent process. pid: 81060, child pid: 81329

  # 親プロセスでの処理
  5.times do
    p('parent')
    sleep(5)
  end

# 子プロセスの終了を待って終了。
#   Process.waitpid(child_pid)
else
  puts "child process. pid: #{Process.pid}"
  # => child process. pid: 81329

  # 子プロセスでの処理
  40.times do
    p('child')
    sleep(5)
  end
end
