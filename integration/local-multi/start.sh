#!/usr/bin/env bash

UPMON="${UPMON:-"./build/upmon"}"

tmux new-session -d -s upmon-local-multi
tmux send "$UPMON --verbose start --config integration/local-multi/config-1.yml" ENTER
tmux split-window -h
tmux send "$UPMON --verbose start --config integration/local-multi/config-2.yml" ENTER
tmux split-window -h
tmux send "$UPMON --verbose start --config integration/local-multi/config-3.yml" ENTER
tmux select-layout tiled
tmux a
