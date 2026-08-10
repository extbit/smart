//
//  Copyright (C) 2012-2025, Duzy Chan <code@extbit.io>, all rights reserverd.
//  Use of this source code is governed by a BSD-style license that can be
//  found in the LICENSE file.
//
//go:build !checkpoints

package smart

import (
	"regexp"
)

const checkpoints = false

func check__string_com(_ Context, _ *compound, _ Value) {}
func check_prefix(_ Context, _ string, _, _ Value, _ *Value) {}
func check_string(_ Context, _ any) func(*string) { return nil }
func check_cache(_ Context, _ any, _ string, _ *valcache, _ []*valcache) {}
func check_uncache(_ Context, _ any, _ string, _ *valcache, _ []*valcache) {}
func check_unmap(_ *uncache_t, _ any, _ *valcache, _ []*valcache) {}
func check_match(_ Context, _, _ Value, _ bool) func(*bool, *Value, *Value, *[]Value) { return nil }
func check_cmp(_ Context, _, _ any) func(*cmpres) { return nil }
func check_com(_ *com_ctx, _, _ []Value, _ *[]Value) {}
func check(_ Context, _, _ Value, _ ...Value) {}
func check_symbolize(Value) func(*[]Symbol) { return nil }
func check_cmp_symbol(_ Context, _, _ Symbol) func(*cmpres) { return nil }
func check_cmp_symbols(_ Context, _, _ []Symbol) func(*cmpres) { return nil }
func check_rule_execute(_ Context, _ *rule, _ []Value) func(*[]Value) { return nil }
func check_traverse_mapped(ctx Context, stems []Value, pat, val Value) {}
func check_evoke_rule_program_res(ctx Context, p *evoke_rule_ctx, pr program_res) {}
func check_rule_executed(ctx Context, x *execution, res Value) {}

func tempdir_check(Context, *project, *def, string) {}
func tempfile_check(Context, *project, Symbol, string, *file) {}

func (*execution) check_interpret(ctx Context, i evaluater, args []Value) func(*Value) { return nil }
func (*exec_buffer) check_line(string, int) {}
func (*exec_ctx) check_exec(i int, src *raw, e error) {}
func (*exec_ctx) run_check(*execution) error { return nil }
func (*exec_ctx) sources_check(Context, int, Value, string) {}
func (*program) execute_check(*execution, *Value) {}
func (*program) execute_check_0(*execution) {}
func (*program) execute_check_1(*execution) {}
func (*scope) check_def(_ Context, _ origin, _ any, _ []Value, _ string) func(**def) { return nil }

func (*modifier_updatefile) check(args []Value, target Value) func(*string, *any) { return nil }

func (*__debug) check() {}
func (*__grep) check(_ *regexp.Regexp, _ string, _, _ Value) {}
func (*__trimprefix) check(_, _, _ Value) {}

func (*compiler) check_ident(*ident_ctx, Context, Value, string, Symbol) {}
func (*compiler) check_sources(Context, Symbol) func(*[]Symbol) { return nil }
func (*compiler) check_assign(ctx Context, id Value, sym Symbol, tok token, vals []Value, d *def, isNew bool, idx int) func(*[]*def) { return nil }
func (*compiler) configure_val_check(_ *execution, _ Symbol, _ Value, _ []Value, _, _ *diagpoint) {}
func (*compiler) check_configure_handle(ctx *execution, op Symbol, vals []Value, a, b *diagpoint) {}
func (*compiler) check_configure_cache(ctx *execution, d *def, isNew, isCached, isDefer bool) {}

func (*dialect_plain) check_evaluate(ctx Context, args, recipes []Value, p *plain) {}
func (*dialect_exec) check_exec_evaluate(ec *exec_ctx) func(*string, *Value) { return nil }
