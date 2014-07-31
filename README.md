cloudgone
=========
ec2 server shutdown management tool

This is a simpke tool that attemps to shutdown inactive ec2 instances
5 minutes before the hour-mark.  This runs on the instance itself,
make sure you launch the instances with `--instance-initiated-shutdown-behavior terminate`
so they are properly terminated.

If you are to run a task on the box for 300 seconds, simply
call `http://localhost:31337/ping/300` from within the box. cloudgone will
extend the lifetime to the next hour-mark if necessary.

## License

http://clkao.mit-license.org/
