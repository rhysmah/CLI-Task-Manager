## CLI Task Manager - Overview
This is a simple CLI task manager that allows you to add, remove, complete, and list tasks. 
Tasks are stored in a boltDB database. 

## Running the Program
To run the task manager, do the following:

1. Open a terminal at the root of the project and build the project via a Makefile:

```bash
make
```

2. Ensure the task manager was properly built by running the following command:

```bash 
./cli-task-manager
```

You should see "Welcome to the CLI Task Manager!" in your terminal.

## Commands
There are four commands: `add`, `remove`, `do` and `list`.

### Adding a Task
```bash
./cli-task-manager add "Task description"
```

### Doing a Task
Completing a task will mark it with `[✓]` when listed. Uncompleted tasks are represented with `[ ]`.

```bash
./cli-task-manager do "Task description"
```

### Removing a Task
Using flags, you can remove a single task or all tasks.

To remove a single task:
```bash
./cli-task-manager remove "Task description"
```

To remove all tasks, you can either use the `-a` or `--all` flag:
```bash
./cli-task-manager remove --all

# OR

./cli-task-manager remove -a
```

### Listing Tasks
Using flags, you can list all tasks, only completed tasks, or only uncompleted tasks.

To list all tasks, don't use a flag:
```bash
./cli-task-manager list
```

To list only completed tasks, use the `-c` or `--completed` flag:
```bash
./cli-task-manager list --completed

# OR

./cli-task-manager list -c
```

To list only uncompleted tasks, use the `-u` or `--uncompleted` flag:
```bash
./cli-task-manager list --uncompleted

# OR

./cli-task-manager list -u
```