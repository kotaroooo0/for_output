# frozen_string_literal: true

if child_pid = fork
  puts "parent process. pid: #{Process.pid}, child pid: #{child_pid}"

  10.times do
    p('parent')
    sleep(1)
  end

# 子プロセスの終了を待って終了。
#   Process.waitpid(child_pid)
else
  puts "child process. pid: #{Process.pid}"

  20.times do
    p('child')
    sleep(1)
  end

  if fork
    p('clild child')
    sleep(3)
  end
end
