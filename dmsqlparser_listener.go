// Code generated from DmSqlParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // DmSqlParser
import "github.com/antlr4-go/antlr/v4"

// DmSqlParserListener is a complete listener for a parse tree produced by DmSqlParser.
type DmSqlParserListener interface {
	antlr.ParseTreeListener

	// EnterDmprogram is called when entering the dmprogram production.
	EnterDmprogram(c *DmprogramContext)

	// EnterSql_clauses is called when entering the sql_clauses production.
	EnterSql_clauses(c *Sql_clausesContext)

	// EnterDdlsql is called when entering the ddlsql production.
	EnterDdlsql(c *DdlsqlContext)

	// EnterDmlsql is called when entering the dmlsql production.
	EnterDmlsql(c *DmlsqlContext)

	// EnterPrivsql is called when entering the privsql production.
	EnterPrivsql(c *PrivsqlContext)

	// EnterOthersql is called when entering the othersql production.
	EnterOthersql(c *OthersqlContext)

	// EnterUtilsql is called when entering the utilsql production.
	EnterUtilsql(c *UtilsqlContext)

	// EnterExplainsql is called when entering the explainsql production.
	EnterExplainsql(c *ExplainsqlContext)

	// EnterShutdown_stmt is called when entering the shutdown_stmt production.
	EnterShutdown_stmt(c *Shutdown_stmtContext)

	// EnterAlter_diskgroup_stmt is called when entering the alter_diskgroup_stmt production.
	EnterAlter_diskgroup_stmt(c *Alter_diskgroup_stmtContext)

	// EnterLocal is called when entering the local production.
	EnterLocal(c *LocalContext)

	// EnterDmsubprogram is called when entering the dmsubprogram production.
	EnterDmsubprogram(c *DmsubprogramContext)

	// EnterDeclare_block is called when entering the declare_block production.
	EnterDeclare_block(c *Declare_blockContext)

	// EnterDecl_var_cur_list_options is called when entering the decl_var_cur_list_options production.
	EnterDecl_var_cur_list_options(c *Decl_var_cur_list_optionsContext)

	// EnterDecl_var_cur_list_1 is called when entering the decl_var_cur_list_1 production.
	EnterDecl_var_cur_list_1(c *Decl_var_cur_list_1Context)

	// EnterDecl_var_cur_list is called when entering the decl_var_cur_list production.
	EnterDecl_var_cur_list(c *Decl_var_cur_listContext)

	// EnterDecl_plsql_type is called when entering the decl_plsql_type production.
	EnterDecl_plsql_type(c *Decl_plsql_typeContext)

	// EnterPlsql_type_def is called when entering the plsql_type_def production.
	EnterPlsql_type_def(c *Plsql_type_defContext)

	// EnterLt_int_lst is called when entering the lt_int_lst production.
	EnterLt_int_lst(c *Lt_int_lstContext)

	// EnterRec_item_def_list is called when entering the rec_item_def_list production.
	EnterRec_item_def_list(c *Rec_item_def_listContext)

	// EnterRec_item_def is called when entering the rec_item_def production.
	EnterRec_item_def(c *Rec_item_defContext)

	// EnterDecl_variable is called when entering the decl_variable production.
	EnterDecl_variable(c *Decl_variableContext)

	// EnterNot_null is called when entering the not_null production.
	EnterNot_null(c *Not_nullContext)

	// EnterPlsql_datatype is called when entering the plsql_datatype production.
	EnterPlsql_datatype(c *Plsql_datatypeContext)

	// EnterDefault_clause_option is called when entering the default_clause_option production.
	EnterDefault_clause_option(c *Default_clause_optionContext)

	// EnterVariable_name_list is called when entering the variable_name_list production.
	EnterVariable_name_list(c *Variable_name_listContext)

	// EnterDecl_except is called when entering the decl_except production.
	EnterDecl_except(c *Decl_exceptContext)

	// EnterPragma_def is called when entering the pragma_def production.
	EnterPragma_def(c *Pragma_defContext)

	// EnterPragma is called when entering the pragma production.
	EnterPragma(c *PragmaContext)

	// EnterPlbody is called when entering the plbody production.
	EnterPlbody(c *PlbodyContext)

	// EnterSs_plbody is called when entering the ss_plbody production.
	EnterSs_plbody(c *Ss_plbodyContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterLabel_list is called when entering the label_list production.
	EnterLabel_list(c *Label_listContext)

	// EnterLabel_list_options is called when entering the label_list_options production.
	EnterLabel_list_options(c *Label_list_optionsContext)

	// EnterLabel_demiliter_l is called when entering the label_demiliter_l production.
	EnterLabel_demiliter_l(c *Label_demiliter_lContext)

	// EnterLabel_demiliter_r is called when entering the label_demiliter_r production.
	EnterLabel_demiliter_r(c *Label_demiliter_rContext)

	// EnterPlsql is called when entering the plsql production.
	EnterPlsql(c *PlsqlContext)

	// EnterUr_option is called when entering the ur_option production.
	EnterUr_option(c *Ur_optionContext)

	// EnterFlashback_trig_enable is called when entering the flashback_trig_enable production.
	EnterFlashback_trig_enable(c *Flashback_trig_enableContext)

	// EnterScn_or_lsn is called when entering the scn_or_lsn production.
	EnterScn_or_lsn(c *Scn_or_lsnContext)

	// EnterFull_table_name_list is called when entering the full_table_name_list production.
	EnterFull_table_name_list(c *Full_table_name_listContext)

	// EnterFlashback_tab_stmt is called when entering the flashback_tab_stmt production.
	EnterFlashback_tab_stmt(c *Flashback_tab_stmtContext)

	// EnterRename is called when entering the rename production.
	EnterRename(c *RenameContext)

	// EnterAlter_system_set_stmt is called when entering the alter_system_set_stmt production.
	EnterAlter_system_set_stmt(c *Alter_system_set_stmtContext)

	// EnterDefer is called when entering the defer production.
	EnterDefer(c *DeferContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// EnterAlter_session_stmt is called when entering the alter_session_stmt production.
	EnterAlter_session_stmt(c *Alter_session_stmtContext)

	// EnterParallel_mode is called when entering the parallel_mode production.
	EnterParallel_mode(c *Parallel_modeContext)

	// EnterParallel_degree is called when entering the parallel_degree production.
	EnterParallel_degree(c *Parallel_degreeContext)

	// EnterPurge is called when entering the purge production.
	EnterPurge(c *PurgeContext)

	// EnterSess_id is called when entering the sess_id production.
	EnterSess_id(c *Sess_idContext)

	// EnterSet_time_zone_string is called when entering the set_time_zone_string production.
	EnterSet_time_zone_string(c *Set_time_zone_stringContext)

	// EnterSess_attr is called when entering the sess_attr production.
	EnterSess_attr(c *Sess_attrContext)

	// EnterSess_attr_val is called when entering the sess_attr_val production.
	EnterSess_attr_val(c *Sess_attr_valContext)

	// EnterSet_schema_stmt is called when entering the set_schema_stmt production.
	EnterSet_schema_stmt(c *Set_schema_stmtContext)

	// EnterPlblock is called when entering the plblock production.
	EnterPlblock(c *PlblockContext)

	// EnterExcept_option is called when entering the except_option production.
	EnterExcept_option(c *Except_optionContext)

	// EnterFinally_option is called when entering the finally_option production.
	EnterFinally_option(c *Finally_optionContext)

	// EnterFinally_tail is called when entering the finally_tail production.
	EnterFinally_tail(c *Finally_tailContext)

	// EnterExcept_handler_list is called when entering the except_handler_list production.
	EnterExcept_handler_list(c *Except_handler_listContext)

	// EnterExcept_handler is called when entering the except_handler production.
	EnterExcept_handler(c *Except_handlerContext)

	// EnterExcept_name is called when entering the except_name production.
	EnterExcept_name(c *Except_nameContext)

	// EnterExcept_list is called when entering the except_list production.
	EnterExcept_list(c *Except_listContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterIf_stmt_clause is called when entering the if_stmt_clause production.
	EnterIf_stmt_clause(c *If_stmt_clauseContext)

	// EnterIf_condition_clause is called when entering the if_condition_clause production.
	EnterIf_condition_clause(c *If_condition_clauseContext)

	// EnterIf_then_clause is called when entering the if_then_clause production.
	EnterIf_then_clause(c *If_then_clauseContext)

	// EnterElseif_lst_option is called when entering the elseif_lst_option production.
	EnterElseif_lst_option(c *Elseif_lst_optionContext)

	// EnterElseif_clause is called when entering the elseif_clause production.
	EnterElseif_clause(c *Elseif_clauseContext)

	// EnterElse_option is called when entering the else_option production.
	EnterElse_option(c *Else_optionContext)

	// EnterSs_if_stmt_clause is called when entering the ss_if_stmt_clause production.
	EnterSs_if_stmt_clause(c *Ss_if_stmt_clauseContext)

	// EnterSs_elseif_lst_option is called when entering the ss_elseif_lst_option production.
	EnterSs_elseif_lst_option(c *Ss_elseif_lst_optionContext)

	// EnterSs_elseif_clause is called when entering the ss_elseif_clause production.
	EnterSs_elseif_clause(c *Ss_elseif_clauseContext)

	// EnterSs_else_option is called when entering the ss_else_option production.
	EnterSs_else_option(c *Ss_else_optionContext)

	// EnterCase_stmt is called when entering the case_stmt production.
	EnterCase_stmt(c *Case_stmtContext)

	// EnterPlsearched_when_clause is called when entering the plsearched_when_clause production.
	EnterPlsearched_when_clause(c *Plsearched_when_clauseContext)

	// EnterPlsearched_when_list is called when entering the plsearched_when_list production.
	EnterPlsearched_when_list(c *Plsearched_when_listContext)

	// EnterCase_option is called when entering the case_option production.
	EnterCase_option(c *Case_optionContext)

	// EnterAssign_stmt is called when entering the assign_stmt production.
	EnterAssign_stmt(c *Assign_stmtContext)

	// EnterAssign_obj is called when entering the assign_obj production.
	EnterAssign_obj(c *Assign_objContext)

	// EnterAssign_obj2 is called when entering the assign_obj2 production.
	EnterAssign_obj2(c *Assign_obj2Context)

	// EnterAssign_op is called when entering the assign_op production.
	EnterAssign_op(c *Assign_opContext)

	// EnterGoto_stmt is called when entering the goto_stmt production.
	EnterGoto_stmt(c *Goto_stmtContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterLoop_stmt is called when entering the loop_stmt production.
	EnterLoop_stmt(c *Loop_stmtContext)

	// EnterRepeat_stmt is called when entering the repeat_stmt production.
	EnterRepeat_stmt(c *Repeat_stmtContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterForall_stmt is called when entering the forall_stmt production.
	EnterForall_stmt(c *Forall_stmtContext)

	// EnterForall_between_option is called when entering the forall_between_option production.
	EnterForall_between_option(c *Forall_between_optionContext)

	// EnterForall_save_exception_option is called when entering the forall_save_exception_option production.
	EnterForall_save_exception_option(c *Forall_save_exception_optionContext)

	// EnterForall_index_values is called when entering the forall_index_values production.
	EnterForall_index_values(c *Forall_index_valuesContext)

	// EnterForall_start is called when entering the forall_start production.
	EnterForall_start(c *Forall_startContext)

	// EnterForall_dml_stmt is called when entering the forall_dml_stmt production.
	EnterForall_dml_stmt(c *Forall_dml_stmtContext)

	// EnterIn_option is called when entering the in_option production.
	EnterIn_option(c *In_optionContext)

	// EnterFor_condition is called when entering the for_condition production.
	EnterFor_condition(c *For_conditionContext)

	// EnterPipe_row_stmt is called when entering the pipe_row_stmt production.
	EnterPipe_row_stmt(c *Pipe_row_stmtContext)

	// EnterExit_stmt is called when entering the exit_stmt production.
	EnterExit_stmt(c *Exit_stmtContext)

	// EnterContinue_stmt is called when entering the continue_stmt production.
	EnterContinue_stmt(c *Continue_stmtContext)

	// EnterNull_stmt is called when entering the null_stmt production.
	EnterNull_stmt(c *Null_stmtContext)

	// EnterPrint_stmt is called when entering the print_stmt production.
	EnterPrint_stmt(c *Print_stmtContext)

	// EnterExecute_stmt is called when entering the execute_stmt production.
	EnterExecute_stmt(c *Execute_stmtContext)

	// EnterDyn_return is called when entering the dyn_return production.
	EnterDyn_return(c *Dyn_returnContext)

	// EnterUsing_clause is called when entering the using_clause production.
	EnterUsing_clause(c *Using_clauseContext)

	// EnterUsing_exp_list is called when entering the using_exp_list production.
	EnterUsing_exp_list(c *Using_exp_listContext)

	// EnterUsing_exp is called when entering the using_exp production.
	EnterUsing_exp(c *Using_expContext)

	// EnterAlter_proc_stmt is called when entering the alter_proc_stmt production.
	EnterAlter_proc_stmt(c *Alter_proc_stmtContext)

	// EnterAlter_func_stmt is called when entering the alter_func_stmt production.
	EnterAlter_func_stmt(c *Alter_func_stmtContext)

	// EnterAlter_package_stmt is called when entering the alter_package_stmt production.
	EnterAlter_package_stmt(c *Alter_package_stmtContext)

	// EnterPkg_type is called when entering the pkg_type production.
	EnterPkg_type(c *Pkg_typeContext)

	// EnterDeclare_opt is called when entering the declare_opt production.
	EnterDeclare_opt(c *Declare_optContext)

	// EnterAlter_table_stmt is called when entering the alter_table_stmt production.
	EnterAlter_table_stmt(c *Alter_table_stmtContext)

	// EnterAlter_tag is called when entering the alter_tag production.
	EnterAlter_tag(c *Alter_tagContext)

	// EnterAlter_index_stmt is called when entering the alter_index_stmt production.
	EnterAlter_index_stmt(c *Alter_index_stmtContext)

	// EnterFull_index_name is called when entering the full_index_name production.
	EnterFull_index_name(c *Full_index_nameContext)

	// EnterAlter_index_action is called when entering the alter_index_action production.
	EnterAlter_index_action(c *Alter_index_actionContext)

	// EnterRebuild_clause is called when entering the rebuild_clause production.
	EnterRebuild_clause(c *Rebuild_clauseContext)

	// EnterExclusive_options is called when entering the exclusive_options production.
	EnterExclusive_options(c *Exclusive_optionsContext)

	// EnterAsynchronous_options is called when entering the asynchronous_options production.
	EnterAsynchronous_options(c *Asynchronous_optionsContext)

	// EnterVisible_clause is called when entering the visible_clause production.
	EnterVisible_clause(c *Visible_clauseContext)

	// EnterColumn_def_list is called when entering the column_def_list production.
	EnterColumn_def_list(c *Column_def_listContext)

	// EnterLock is called when entering the lock production.
	EnterLock(c *LockContext)

	// EnterAlter_table_partition_action is called when entering the alter_table_partition_action production.
	EnterAlter_table_partition_action(c *Alter_table_partition_actionContext)

	// EnterTemplate_info is called when entering the template_info production.
	EnterTemplate_info(c *Template_infoContext)

	// EnterTemplate_item_2 is called when entering the template_item_2 production.
	EnterTemplate_item_2(c *Template_item_2Context)

	// EnterTemplate_item_1 is called when entering the template_item_1 production.
	EnterTemplate_item_1(c *Template_item_1Context)

	// EnterIncluding_indexes is called when entering the including_indexes production.
	EnterIncluding_indexes(c *Including_indexesContext)

	// EnterTruncate_partition_name is called when entering the truncate_partition_name production.
	EnterTruncate_partition_name(c *Truncate_partition_nameContext)

	// EnterCons_enable is called when entering the cons_enable production.
	EnterCons_enable(c *Cons_enableContext)

	// EnterReuse_storage_option is called when entering the reuse_storage_option production.
	EnterReuse_storage_option(c *Reuse_storage_optionContext)

	// EnterAlter_table_action is called when entering the alter_table_action production.
	EnterAlter_table_action(c *Alter_table_actionContext)

	// EnterFast_flag is called when entering the fast_flag production.
	EnterFast_flag(c *Fast_flagContext)

	// EnterStorage_stat_flag is called when entering the storage_stat_flag production.
	EnterStorage_stat_flag(c *Storage_stat_flagContext)

	// EnterStorage_stat_cols is called when entering the storage_stat_cols production.
	EnterStorage_stat_cols(c *Storage_stat_colsContext)

	// EnterHfs_rebuild_level is called when entering the hfs_rebuild_level production.
	EnterHfs_rebuild_level(c *Hfs_rebuild_levelContext)

	// EnterAta_lock_option is called when entering the ata_lock_option production.
	EnterAta_lock_option(c *Ata_lock_optionContext)

	// EnterMdf_column_def_list is called when entering the mdf_column_def_list production.
	EnterMdf_column_def_list(c *Mdf_column_def_listContext)

	// EnterMdf_column_def is called when entering the mdf_column_def production.
	EnterMdf_column_def(c *Mdf_column_defContext)

	// EnterColumn_def is called when entering the column_def production.
	EnterColumn_def(c *Column_defContext)

	// EnterColumn_def_ex is called when entering the column_def_ex production.
	EnterColumn_def_ex(c *Column_def_exContext)

	// EnterColumn_def_low is called when entering the column_def_low production.
	EnterColumn_def_low(c *Column_def_lowContext)

	// EnterVirtual_column_datatype is called when entering the virtual_column_datatype production.
	EnterVirtual_column_datatype(c *Virtual_column_datatypeContext)

	// EnterVirtual_column_generated is called when entering the virtual_column_generated production.
	EnterVirtual_column_generated(c *Virtual_column_generatedContext)

	// EnterVirtual_column_virtual is called when entering the virtual_column_virtual production.
	EnterVirtual_column_virtual(c *Virtual_column_virtualContext)

	// EnterVirtual_column_visible is called when entering the virtual_column_visible production.
	EnterVirtual_column_visible(c *Virtual_column_visibleContext)

	// EnterVirtual_column_def is called when entering the virtual_column_def production.
	EnterVirtual_column_def(c *Virtual_column_defContext)

	// EnterCharset_option is called when entering the charset_option production.
	EnterCharset_option(c *Charset_optionContext)

	// EnterColumn_def_4_option is called when entering the column_def_4_option production.
	EnterColumn_def_4_option(c *Column_def_4_optionContext)

	// EnterAuto_update_clause is called when entering the auto_update_clause production.
	EnterAuto_update_clause(c *Auto_update_clauseContext)

	// EnterUpdate_exp is called when entering the update_exp production.
	EnterUpdate_exp(c *Update_expContext)

	// EnterIdentity_clause is called when entering the identity_clause production.
	EnterIdentity_clause(c *Identity_clauseContext)

	// EnterDefault_clause_with_on_null_opt is called when entering the default_clause_with_on_null_opt production.
	EnterDefault_clause_with_on_null_opt(c *Default_clause_with_on_null_optContext)

	// EnterDefault_clause is called when entering the default_clause production.
	EnterDefault_clause(c *Default_clauseContext)

	// EnterDefault_exp is called when entering the default_exp production.
	EnterDefault_exp(c *Default_expContext)

	// EnterColumn_constraint_def is called when entering the column_constraint_def production.
	EnterColumn_constraint_def(c *Column_constraint_defContext)

	// EnterConstraint_name_def_options is called when entering the constraint_name_def_options production.
	EnterConstraint_name_def_options(c *Constraint_name_def_optionsContext)

	// EnterConstraint_name_def is called when entering the constraint_name_def production.
	EnterConstraint_name_def(c *Constraint_name_defContext)

	// EnterColumn_constraints is called when entering the column_constraints production.
	EnterColumn_constraints(c *Column_constraintsContext)

	// EnterColumn_constraint is called when entering the column_constraint production.
	EnterColumn_constraint(c *Column_constraintContext)

	// EnterColumn_constraint_action is called when entering the column_constraint_action production.
	EnterColumn_constraint_action(c *Column_constraint_actionContext)

	// EnterNot_null_spec is called when entering the not_null_spec production.
	EnterNot_null_spec(c *Not_null_specContext)

	// EnterUnique_spec is called when entering the unique_spec production.
	EnterUnique_spec(c *Unique_specContext)

	// EnterRefs_spec is called when entering the refs_spec production.
	EnterRefs_spec(c *Refs_specContext)

	// EnterRefs_spec_action is called when entering the refs_spec_action production.
	EnterRefs_spec_action(c *Refs_spec_actionContext)

	// EnterForeign_key is called when entering the foreign_key production.
	EnterForeign_key(c *Foreign_keyContext)

	// EnterRefd_table_and_columns is called when entering the refd_table_and_columns production.
	EnterRefd_table_and_columns(c *Refd_table_and_columnsContext)

	// EnterRef_column_list is called when entering the ref_column_list production.
	EnterRef_column_list(c *Ref_column_listContext)

	// EnterColumn_list is called when entering the column_list production.
	EnterColumn_list(c *Column_listContext)

	// EnterColumn_list2 is called when entering the column_list2 production.
	EnterColumn_list2(c *Column_list2Context)

	// EnterFull_column_list is called when entering the full_column_list production.
	EnterFull_column_list(c *Full_column_listContext)

	// EnterColumn_list_list is called when entering the column_list_list production.
	EnterColumn_list_list(c *Column_list_listContext)

	// EnterDrop_column_list is called when entering the drop_column_list production.
	EnterDrop_column_list(c *Drop_column_listContext)

	// EnterMatch_option is called when entering the match_option production.
	EnterMatch_option(c *Match_optionContext)

	// EnterMatch_type is called when entering the match_type production.
	EnterMatch_type(c *Match_typeContext)

	// EnterRef_triggered_action is called when entering the ref_triggered_action production.
	EnterRef_triggered_action(c *Ref_triggered_actionContext)

	// EnterUpdate_rule is called when entering the update_rule production.
	EnterUpdate_rule(c *Update_ruleContext)

	// EnterDelete_rule is called when entering the delete_rule production.
	EnterDelete_rule(c *Delete_ruleContext)

	// EnterRef_action is called when entering the ref_action production.
	EnterRef_action(c *Ref_actionContext)

	// EnterCheck_constraint_def is called when entering the check_constraint_def production.
	EnterCheck_constraint_def(c *Check_constraint_defContext)

	// EnterCheck_condition is called when entering the check_condition production.
	EnterCheck_condition(c *Check_conditionContext)

	// EnterRestrict_cascade is called when entering the restrict_cascade production.
	EnterRestrict_cascade(c *Restrict_cascadeContext)

	// EnterCascade_opt is called when entering the cascade_opt production.
	EnterCascade_opt(c *Cascade_optContext)

	// EnterConstraint_name_options is called when entering the constraint_name_options production.
	EnterConstraint_name_options(c *Constraint_name_optionsContext)

	// EnterCheck_option_def_true is called when entering the check_option_def_true production.
	EnterCheck_option_def_true(c *Check_option_def_trueContext)

	// EnterConstraint_attributes_options is called when entering the constraint_attributes_options production.
	EnterConstraint_attributes_options(c *Constraint_attributes_optionsContext)

	// EnterConstraint_attributes is called when entering the constraint_attributes production.
	EnterConstraint_attributes(c *Constraint_attributesContext)

	// EnterDeferrable_option is called when entering the deferrable_option production.
	EnterDeferrable_option(c *Deferrable_optionContext)

	// EnterConstraint_check_time is called when entering the constraint_check_time production.
	EnterConstraint_check_time(c *Constraint_check_timeContext)

	// EnterTable_constraint_clause is called when entering the table_constraint_clause production.
	EnterTable_constraint_clause(c *Table_constraint_clauseContext)

	// EnterTable_constraint is called when entering the table_constraint production.
	EnterTable_constraint(c *Table_constraintContext)

	// EnterUsing_index_clause is called when entering the using_index_clause production.
	EnterUsing_index_clause(c *Using_index_clauseContext)

	// EnterForeign_key_clause is called when entering the foreign_key_clause production.
	EnterForeign_key_clause(c *Foreign_key_clauseContext)

	// EnterAlter_trigger_stmt is called when entering the alter_trigger_stmt production.
	EnterAlter_trigger_stmt(c *Alter_trigger_stmtContext)

	// EnterAlter_trigger_option is called when entering the alter_trigger_option production.
	EnterAlter_trigger_option(c *Alter_trigger_optionContext)

	// EnterAlter_table_partition_action_options is called when entering the alter_table_partition_action_options production.
	EnterAlter_table_partition_action_options(c *Alter_table_partition_action_optionsContext)

	// EnterRefresh_materialized_view_stmt is called when entering the refresh_materialized_view_stmt production.
	EnterRefresh_materialized_view_stmt(c *Refresh_materialized_view_stmtContext)

	// EnterComplete_del_null is called when entering the complete_del_null production.
	EnterComplete_del_null(c *Complete_del_nullContext)

	// EnterRefresh_complete_del is called when entering the refresh_complete_del production.
	EnterRefresh_complete_del(c *Refresh_complete_delContext)

	// EnterAlter_materialized_view_stmt is called when entering the alter_materialized_view_stmt production.
	EnterAlter_materialized_view_stmt(c *Alter_materialized_view_stmtContext)

	// EnterAlter_view_stmt is called when entering the alter_view_stmt production.
	EnterAlter_view_stmt(c *Alter_view_stmtContext)

	// EnterAlter_view_action is called when entering the alter_view_action production.
	EnterAlter_view_action(c *Alter_view_actionContext)

	// EnterCons_novalidate is called when entering the cons_novalidate production.
	EnterCons_novalidate(c *Cons_novalidateContext)

	// EnterView_constraint_clause is called when entering the view_constraint_clause production.
	EnterView_constraint_clause(c *View_constraint_clauseContext)

	// EnterView_constraint is called when entering the view_constraint production.
	EnterView_constraint(c *View_constraintContext)

	// EnterView_unique_spec is called when entering the view_unique_spec production.
	EnterView_unique_spec(c *View_unique_specContext)

	// EnterView_refs_spec is called when entering the view_refs_spec production.
	EnterView_refs_spec(c *View_refs_specContext)

	// EnterView_refs_spec_action is called when entering the view_refs_spec_action production.
	EnterView_refs_spec_action(c *View_refs_spec_actionContext)

	// EnterCall_proc_stmt is called when entering the call_proc_stmt production.
	EnterCall_proc_stmt(c *Call_proc_stmtContext)

	// EnterRaw_call_proc_stmt is called when entering the raw_call_proc_stmt production.
	EnterRaw_call_proc_stmt(c *Raw_call_proc_stmtContext)

	// EnterCall_proc_stmt_2 is called when entering the call_proc_stmt_2 production.
	EnterCall_proc_stmt_2(c *Call_proc_stmt_2Context)

	// EnterExec_proc_stmt is called when entering the exec_proc_stmt production.
	EnterExec_proc_stmt(c *Exec_proc_stmtContext)

	// EnterDblink_clause is called when entering the dblink_clause production.
	EnterDblink_clause(c *Dblink_clauseContext)

	// EnterDblink_clause2 is called when entering the dblink_clause2 production.
	EnterDblink_clause2(c *Dblink_clause2Context)

	// EnterParam_list is called when entering the param_list production.
	EnterParam_list(c *Param_listContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterRaw_exp_list is called when entering the raw_exp_list production.
	EnterRaw_exp_list(c *Raw_exp_listContext)

	// EnterExp_list_2 is called when entering the exp_list_2 production.
	EnterExp_list_2(c *Exp_list_2Context)

	// EnterExp_list is called when entering the exp_list production.
	EnterExp_list(c *Exp_listContext)

	// EnterIns_exp_list is called when entering the ins_exp_list production.
	EnterIns_exp_list(c *Ins_exp_listContext)

	// EnterLt_exp is called when entering the lt_exp production.
	EnterLt_exp(c *Lt_expContext)

	// EnterRange_partition_exp is called when entering the range_partition_exp production.
	EnterRange_partition_exp(c *Range_partition_expContext)

	// EnterRange_partition_exp_list is called when entering the range_partition_exp_list production.
	EnterRange_partition_exp_list(c *Range_partition_exp_listContext)

	// EnterList_partition_exp is called when entering the list_partition_exp production.
	EnterList_partition_exp(c *List_partition_expContext)

	// EnterList_partition_exp_list is called when entering the list_partition_exp_list production.
	EnterList_partition_exp_list(c *List_partition_exp_listContext)

	// EnterList_partition_value_list is called when entering the list_partition_value_list production.
	EnterList_partition_value_list(c *List_partition_value_listContext)

	// EnterClose_cursor_stmt is called when entering the close_cursor_stmt production.
	EnterClose_cursor_stmt(c *Close_cursor_stmtContext)

	// EnterClose_cursor_statement is called when entering the close_cursor_statement production.
	EnterClose_cursor_statement(c *Close_cursor_statementContext)

	// EnterBegin_trans_stmt is called when entering the begin_trans_stmt production.
	EnterBegin_trans_stmt(c *Begin_trans_stmtContext)

	// EnterCommit_trans_stmt is called when entering the commit_trans_stmt production.
	EnterCommit_trans_stmt(c *Commit_trans_stmtContext)

	// EnterCommit_head is called when entering the commit_head production.
	EnterCommit_head(c *Commit_headContext)

	// EnterCommit_tail is called when entering the commit_tail production.
	EnterCommit_tail(c *Commit_tailContext)

	// EnterCommit_wait_immed_option is called when entering the commit_wait_immed_option production.
	EnterCommit_wait_immed_option(c *Commit_wait_immed_optionContext)

	// EnterConnect_stmt is called when entering the connect_stmt production.
	EnterConnect_stmt(c *Connect_stmtContext)

	// EnterPassword is called when entering the password production.
	EnterPassword(c *PasswordContext)

	// EnterTs_storage is called when entering the ts_storage production.
	EnterTs_storage(c *Ts_storageContext)

	// EnterTs_storage_clause is called when entering the ts_storage_clause production.
	EnterTs_storage_clause(c *Ts_storage_clauseContext)

	// EnterCreate_tablespace_stmt is called when entering the create_tablespace_stmt production.
	EnterCreate_tablespace_stmt(c *Create_tablespace_stmtContext)

	// EnterCtss_with_clause is called when entering the ctss_with_clause production.
	EnterCtss_with_clause(c *Ctss_with_clauseContext)

	// EnterCreate_tablespace_set_stmt is called when entering the create_tablespace_set_stmt production.
	EnterCreate_tablespace_set_stmt(c *Create_tablespace_set_stmtContext)

	// EnterAlter_tablespace_set_stmt is called when entering the alter_tablespace_set_stmt production.
	EnterAlter_tablespace_set_stmt(c *Alter_tablespace_set_stmtContext)

	// EnterCache is called when entering the cache production.
	EnterCache(c *CacheContext)

	// EnterAlter_tablespace_stmt is called when entering the alter_tablespace_stmt production.
	EnterAlter_tablespace_stmt(c *Alter_tablespace_stmtContext)

	// EnterKeep is called when entering the keep production.
	EnterKeep(c *KeepContext)

	// EnterAlter_tablespace_action is called when entering the alter_tablespace_action production.
	EnterAlter_tablespace_action(c *Alter_tablespace_actionContext)

	// EnterFile_list is called when entering the file_list production.
	EnterFile_list(c *File_listContext)

	// EnterPathname_list is called when entering the pathname_list production.
	EnterPathname_list(c *Pathname_listContext)

	// EnterInteger_list is called when entering the integer_list production.
	EnterInteger_list(c *Integer_listContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterMirror is called when entering the mirror production.
	EnterMirror(c *MirrorContext)

	// EnterAutoextend_nextsize is called when entering the autoextend_nextsize production.
	EnterAutoextend_nextsize(c *Autoextend_nextsizeContext)

	// EnterAutoextend_maxsize is called when entering the autoextend_maxsize production.
	EnterAutoextend_maxsize(c *Autoextend_maxsizeContext)

	// EnterAutoextend is called when entering the autoextend production.
	EnterAutoextend(c *AutoextendContext)

	// EnterOn_raft is called when entering the on_raft production.
	EnterOn_raft(c *On_raftContext)

	// EnterArchfile is called when entering the archfile production.
	EnterArchfile(c *ArchfileContext)

	// EnterArchflag is called when entering the archflag production.
	EnterArchflag(c *ArchflagContext)

	// EnterArchstyle_options is called when entering the archstyle_options production.
	EnterArchstyle_options(c *Archstyle_optionsContext)

	// EnterArchstyle is called when entering the archstyle production.
	EnterArchstyle(c *ArchstyleContext)

	// EnterArchdir is called when entering the archdir production.
	EnterArchdir(c *ArchdirContext)

	// EnterBakfile is called when entering the bakfile production.
	EnterBakfile(c *BakfileContext)

	// EnterParameters_option_list is called when entering the parameters_option_list production.
	EnterParameters_option_list(c *Parameters_option_listContext)

	// EnterParameter_option_list is called when entering the parameter_option_list production.
	EnterParameter_option_list(c *Parameter_option_listContext)

	// EnterParameter_option is called when entering the parameter_option production.
	EnterParameter_option(c *Parameter_optionContext)

	// EnterPathname is called when entering the pathname production.
	EnterPathname(c *PathnameContext)

	// EnterPathname_options is called when entering the pathname_options production.
	EnterPathname_options(c *Pathname_optionsContext)

	// EnterBackup_stmt is called when entering the backup_stmt production.
	EnterBackup_stmt(c *Backup_stmtContext)

	// EnterBack_range_option is called when entering the back_range_option production.
	EnterBack_range_option(c *Back_range_optionContext)

	// EnterBack_archive_spec_null is called when entering the back_archive_spec_null production.
	EnterBack_archive_spec_null(c *Back_archive_spec_nullContext)

	// EnterNot_backed_up is called when entering the not_backed_up production.
	EnterNot_backed_up(c *Not_backed_upContext)

	// EnterArchive_spec is called when entering the archive_spec production.
	EnterArchive_spec(c *Archive_specContext)

	// EnterSpec_lsn is called when entering the spec_lsn production.
	EnterSpec_lsn(c *Spec_lsnContext)

	// EnterBackup_delete_archive is called when entering the backup_delete_archive production.
	EnterBackup_delete_archive(c *Backup_delete_archiveContext)

	// EnterBackup_options is called when entering the backup_options production.
	EnterBackup_options(c *Backup_optionsContext)

	// EnterCumulative is called when entering the cumulative production.
	EnterCumulative(c *CumulativeContext)

	// EnterWith_bak_dir_list is called when entering the with_bak_dir_list production.
	EnterWith_bak_dir_list(c *With_bak_dir_listContext)

	// EnterBase_on_backup is called when entering the base_on_backup production.
	EnterBase_on_backup(c *Base_on_backupContext)

	// EnterBackup_to_options is called when entering the backup_to_options production.
	EnterBackup_to_options(c *Backup_to_optionsContext)

	// EnterBackup_path_null is called when entering the backup_path_null production.
	EnterBackup_path_null(c *Backup_path_nullContext)

	// EnterDevice_type is called when entering the device_type production.
	EnterDevice_type(c *Device_typeContext)

	// EnterParms_command is called when entering the parms_command production.
	EnterParms_command(c *Parms_commandContext)

	// EnterMedia_name is called when entering the media_name production.
	EnterMedia_name(c *Media_nameContext)

	// EnterBackup_desc_options is called when entering the backup_desc_options production.
	EnterBackup_desc_options(c *Backup_desc_optionsContext)

	// EnterBackup_desc is called when entering the backup_desc production.
	EnterBackup_desc(c *Backup_descContext)

	// EnterBackup_maxsize is called when entering the backup_maxsize production.
	EnterBackup_maxsize(c *Backup_maxsizeContext)

	// EnterBackup_limit is called when entering the backup_limit production.
	EnterBackup_limit(c *Backup_limitContext)

	// EnterBackup_identified is called when entering the backup_identified production.
	EnterBackup_identified(c *Backup_identifiedContext)

	// EnterBackup_compressed is called when entering the backup_compressed production.
	EnterBackup_compressed(c *Backup_compressedContext)

	// EnterBackup_without is called when entering the backup_without production.
	EnterBackup_without(c *Backup_withoutContext)

	// EnterBackup_tsk_thread_num_null is called when entering the backup_tsk_thread_num_null production.
	EnterBackup_tsk_thread_num_null(c *Backup_tsk_thread_num_nullContext)

	// EnterBackup_parallel_dir is called when entering the backup_parallel_dir production.
	EnterBackup_parallel_dir(c *Backup_parallel_dirContext)

	// EnterBackup_trace_file_level is called when entering the backup_trace_file_level production.
	EnterBackup_trace_file_level(c *Backup_trace_file_levelContext)

	// EnterRestore_stmt is called when entering the restore_stmt production.
	EnterRestore_stmt(c *Restore_stmtContext)

	// EnterRestore_datafile_lst is called when entering the restore_datafile_lst production.
	EnterRestore_datafile_lst(c *Restore_datafile_lstContext)

	// EnterRestore_mapped_file is called when entering the restore_mapped_file production.
	EnterRestore_mapped_file(c *Restore_mapped_fileContext)

	// EnterMapped_file is called when entering the mapped_file production.
	EnterMapped_file(c *Mapped_fileContext)

	// EnterRes_struct is called when entering the res_struct production.
	EnterRes_struct(c *Res_structContext)

	// EnterTsk_thread_num is called when entering the tsk_thread_num production.
	EnterTsk_thread_num(c *Tsk_thread_numContext)

	// EnterRestore_tsk_thread_num_null is called when entering the restore_tsk_thread_num_null production.
	EnterRestore_tsk_thread_num_null(c *Restore_tsk_thread_num_nullContext)

	// EnterRestore_parallel is called when entering the restore_parallel production.
	EnterRestore_parallel(c *Restore_parallelContext)

	// EnterFull_table_name_options is called when entering the full_table_name_options production.
	EnterFull_table_name_options(c *Full_table_name_optionsContext)

	// EnterRes_without_index_constraint is called when entering the res_without_index_constraint production.
	EnterRes_without_index_constraint(c *Res_without_index_constraintContext)

	// EnterRestore_from is called when entering the restore_from production.
	EnterRestore_from(c *Restore_fromContext)

	// EnterRes_until is called when entering the res_until production.
	EnterRes_until(c *Res_untilContext)

	// EnterRestore_file_list_options is called when entering the restore_file_list_options production.
	EnterRestore_file_list_options(c *Restore_file_list_optionsContext)

	// EnterMirror_file_list_options is called when entering the mirror_file_list_options production.
	EnterMirror_file_list_options(c *Mirror_file_list_optionsContext)

	// EnterRestore_trace_file_level is called when entering the restore_trace_file_level production.
	EnterRestore_trace_file_level(c *Restore_trace_file_levelContext)

	// EnterRestore_file_list is called when entering the restore_file_list production.
	EnterRestore_file_list(c *Restore_file_listContext)

	// EnterRestore_file is called when entering the restore_file production.
	EnterRestore_file(c *Restore_fileContext)

	// EnterMirror_file_list is called when entering the mirror_file_list production.
	EnterMirror_file_list(c *Mirror_file_listContext)

	// EnterMirror_file is called when entering the mirror_file production.
	EnterMirror_file(c *Mirror_fileContext)

	// EnterWith_bak_arch_dir_list is called when entering the with_bak_arch_dir_list production.
	EnterWith_bak_arch_dir_list(c *With_bak_arch_dir_listContext)

	// EnterRestore_identified is called when entering the restore_identified production.
	EnterRestore_identified(c *Restore_identifiedContext)

	// EnterCreate_func_stmt is called when entering the create_func_stmt production.
	EnterCreate_func_stmt(c *Create_func_stmtContext)

	// EnterFunc_aggr_clause is called when entering the func_aggr_clause production.
	EnterFunc_aggr_clause(c *Func_aggr_clauseContext)

	// EnterPipelined_options is called when entering the pipelined_options production.
	EnterPipelined_options(c *Pipelined_optionsContext)

	// EnterReplace_option is called when entering the replace_option production.
	EnterReplace_option(c *Replace_optionContext)

	// EnterEdit_options is called when entering the edit_options production.
	EnterEdit_options(c *Edit_optionsContext)

	// EnterEncryption_option is called when entering the encryption_option production.
	EnterEncryption_option(c *Encryption_optionContext)

	// EnterCalc_option is called when entering the calc_option production.
	EnterCalc_option(c *Calc_optionContext)

	// EnterFunc_action is called when entering the func_action production.
	EnterFunc_action(c *Func_actionContext)

	// EnterFunc_call_options is called when entering the func_call_options production.
	EnterFunc_call_options(c *Func_call_optionsContext)

	// EnterFunc_call_option_list is called when entering the func_call_option_list production.
	EnterFunc_call_option_list(c *Func_call_option_listContext)

	// EnterFunc_call_option is called when entering the func_call_option production.
	EnterFunc_call_option(c *Func_call_optionContext)

	// EnterInvoker_rights_clause_options is called when entering the invoker_rights_clause_options production.
	EnterInvoker_rights_clause_options(c *Invoker_rights_clause_optionsContext)

	// EnterInvoker_rights_clause is called when entering the invoker_rights_clause production.
	EnterInvoker_rights_clause(c *Invoker_rights_clauseContext)

	// EnterDeterministic_clause_options is called when entering the deterministic_clause_options production.
	EnterDeterministic_clause_options(c *Deterministic_clause_optionsContext)

	// EnterDeterministic_clause is called when entering the deterministic_clause production.
	EnterDeterministic_clause(c *Deterministic_clauseContext)

	// EnterFunc_call_option2_options is called when entering the func_call_option2_options production.
	EnterFunc_call_option2_options(c *Func_call_option2_optionsContext)

	// EnterFunc_call_option_list2 is called when entering the func_call_option_list2 production.
	EnterFunc_call_option_list2(c *Func_call_option_list2Context)

	// EnterFunc_call_option2 is called when entering the func_call_option2 production.
	EnterFunc_call_option2(c *Func_call_option2Context)

	// EnterResult_cache_clause is called when entering the result_cache_clause production.
	EnterResult_cache_clause(c *Result_cache_clauseContext)

	// EnterInner_fun_name is called when entering the inner_fun_name production.
	EnterInner_fun_name(c *Inner_fun_nameContext)

	// EnterPlatform_type is called when entering the platform_type production.
	EnterPlatform_type(c *Platform_typeContext)

	// EnterParam_def_list_option is called when entering the param_def_list_option production.
	EnterParam_def_list_option(c *Param_def_list_optionContext)

	// EnterParam_def_list is called when entering the param_def_list production.
	EnterParam_def_list(c *Param_def_listContext)

	// EnterParam_def is called when entering the param_def production.
	EnterParam_def(c *Param_defContext)

	// EnterParam_in_out_option is called when entering the param_in_out_option production.
	EnterParam_in_out_option(c *Param_in_out_optionContext)

	// EnterIs_as is called when entering the is_as production.
	EnterIs_as(c *Is_asContext)

	// EnterStat_on_org_stmt is called when entering the stat_on_org_stmt production.
	EnterStat_on_org_stmt(c *Stat_on_org_stmtContext)

	// EnterStat_size is called when entering the stat_size production.
	EnterStat_size(c *Stat_sizeContext)

	// EnterStat_para is called when entering the stat_para production.
	EnterStat_para(c *Stat_paraContext)

	// EnterStat_summarize is called when entering the stat_summarize production.
	EnterStat_summarize(c *Stat_summarizeContext)

	// EnterMstat_ex is called when entering the mstat_ex production.
	EnterMstat_ex(c *Mstat_exContext)

	// EnterIndexid is called when entering the indexid production.
	EnterIndexid(c *IndexidContext)

	// EnterGlobal_tag is called when entering the global_tag production.
	EnterGlobal_tag(c *Global_tagContext)

	// EnterBm_join_index_clause is called when entering the bm_join_index_clause production.
	EnterBm_join_index_clause(c *Bm_join_index_clauseContext)

	// EnterParallel_stmt is called when entering the parallel_stmt production.
	EnterParallel_stmt(c *Parallel_stmtContext)

	// EnterCreate_index_stmt is called when entering the create_index_stmt production.
	EnterCreate_index_stmt(c *Create_index_stmtContext)

	// EnterWith_inner is called when entering the with_inner production.
	EnterWith_inner(c *With_innerContext)

	// EnterIndex_no_sort is called when entering the index_no_sort production.
	EnterIndex_no_sort(c *Index_no_sortContext)

	// EnterOnline_options is called when entering the online_options production.
	EnterOnline_options(c *Online_optionsContext)

	// EnterUnusable_options is called when entering the unusable_options production.
	EnterUnusable_options(c *Unusable_optionsContext)

	// EnterReverse_options is called when entering the reverse_options production.
	EnterReverse_options(c *Reverse_optionsContext)

	// EnterIndex_column_list is called when entering the index_column_list production.
	EnterIndex_column_list(c *Index_column_listContext)

	// EnterIndex_column_name is called when entering the index_column_name production.
	EnterIndex_column_name(c *Index_column_nameContext)

	// EnterStorage_hash_tag is called when entering the storage_hash_tag production.
	EnterStorage_hash_tag(c *Storage_hash_tagContext)

	// EnterStorage_hash is called when entering the storage_hash production.
	EnterStorage_hash(c *Storage_hashContext)

	// EnterStorage_tag is called when entering the storage_tag production.
	EnterStorage_tag(c *Storage_tagContext)

	// EnterStorage_tag_nn is called when entering the storage_tag_nn production.
	EnterStorage_tag_nn(c *Storage_tag_nnContext)

	// EnterTablespace_clause is called when entering the tablespace_clause production.
	EnterTablespace_clause(c *Tablespace_clauseContext)

	// EnterObject_table_substitution_clause is called when entering the object_table_substitution_clause production.
	EnterObject_table_substitution_clause(c *Object_table_substitution_clauseContext)

	// EnterObject_table_substitution is called when entering the object_table_substitution production.
	EnterObject_table_substitution(c *Object_table_substitutionContext)

	// EnterOid_clause is called when entering the oid_clause production.
	EnterOid_clause(c *Oid_clauseContext)

	// EnterOid_gen_type is called when entering the oid_gen_type production.
	EnterOid_gen_type(c *Oid_gen_typeContext)

	// EnterOid_index_clause is called when entering the oid_index_clause production.
	EnterOid_index_clause(c *Oid_index_clauseContext)

	// EnterOid_tablespace_clause is called when entering the oid_tablespace_clause production.
	EnterOid_tablespace_clause(c *Oid_tablespace_clauseContext)

	// EnterOid_tablespace_name is called when entering the oid_tablespace_name production.
	EnterOid_tablespace_name(c *Oid_tablespace_nameContext)

	// EnterLocal_option is called when entering the local_option production.
	EnterLocal_option(c *Local_optionContext)

	// EnterStorage_list is called when entering the storage_list production.
	EnterStorage_list(c *Storage_listContext)

	// EnterStorage_hashpartmap is called when entering the storage_hashpartmap production.
	EnterStorage_hashpartmap(c *Storage_hashpartmapContext)

	// EnterStorage is called when entering the storage production.
	EnterStorage(c *StorageContext)

	// EnterId_list is called when entering the id_list production.
	EnterId_list(c *Id_listContext)

	// EnterCreate_proc_stmt is called when entering the create_proc_stmt production.
	EnterCreate_proc_stmt(c *Create_proc_stmtContext)

	// EnterCreate_package_stmt is called when entering the create_package_stmt production.
	EnterCreate_package_stmt(c *Create_package_stmtContext)

	// EnterPkg_cls_flag is called when entering the pkg_cls_flag production.
	EnterPkg_cls_flag(c *Pkg_cls_flagContext)

	// EnterBlk_end_option is called when entering the blk_end_option production.
	EnterBlk_end_option(c *Blk_end_optionContext)

	// EnterBlk_end_option_low is called when entering the blk_end_option_low production.
	EnterBlk_end_option_low(c *Blk_end_option_lowContext)

	// EnterPackage_def_list_options is called when entering the package_def_list_options production.
	EnterPackage_def_list_options(c *Package_def_list_optionsContext)

	// EnterPackage_def_list is called when entering the package_def_list production.
	EnterPackage_def_list(c *Package_def_listContext)

	// EnterPackage_def is called when entering the package_def production.
	EnterPackage_def(c *Package_defContext)

	// EnterRestrict_param_lst is called when entering the restrict_param_lst production.
	EnterRestrict_param_lst(c *Restrict_param_lstContext)

	// EnterCreate_package_body_stmt is called when entering the create_package_body_stmt production.
	EnterCreate_package_body_stmt(c *Create_package_body_stmtContext)

	// EnterCreate_pkg_body_header is called when entering the create_pkg_body_header production.
	EnterCreate_pkg_body_header(c *Create_pkg_body_headerContext)

	// EnterPkg_cls_body_flag is called when entering the pkg_cls_body_flag production.
	EnterPkg_cls_body_flag(c *Pkg_cls_body_flagContext)

	// EnterPackage_body_init_option is called when entering the package_body_init_option production.
	EnterPackage_body_init_option(c *Package_body_init_optionContext)

	// EnterPackage_body_def_list is called when entering the package_body_def_list production.
	EnterPackage_body_def_list(c *Package_body_def_listContext)

	// EnterPackage_body_def is called when entering the package_body_def production.
	EnterPackage_body_def(c *Package_body_defContext)

	// EnterPackage_body_def2 is called when entering the package_body_def2 production.
	EnterPackage_body_def2(c *Package_body_def2Context)

	// EnterCheck_exec_options is called when entering the check_exec_options production.
	EnterCheck_exec_options(c *Check_exec_optionsContext)

	// EnterSubpg_decl_stmt is called when entering the subpg_decl_stmt production.
	EnterSubpg_decl_stmt(c *Subpg_decl_stmtContext)

	// EnterCreate_type_stmt is called when entering the create_type_stmt production.
	EnterCreate_type_stmt(c *Create_type_stmtContext)

	// EnterForce_option is called when entering the force_option production.
	EnterForce_option(c *Force_optionContext)

	// EnterObject_option is called when entering the object_option production.
	EnterObject_option(c *Object_optionContext)

	// EnterUnder_option is called when entering the under_option production.
	EnterUnder_option(c *Under_optionContext)

	// EnterObject_def_list_options is called when entering the object_def_list_options production.
	EnterObject_def_list_options(c *Object_def_list_optionsContext)

	// EnterObject_def_list is called when entering the object_def_list production.
	EnterObject_def_list(c *Object_def_listContext)

	// EnterObject_def is called when entering the object_def production.
	EnterObject_def(c *Object_defContext)

	// EnterMember_static is called when entering the member_static production.
	EnterMember_static(c *Member_staticContext)

	// EnterMethod_inherit_options is called when entering the method_inherit_options production.
	EnterMethod_inherit_options(c *Method_inherit_optionsContext)

	// EnterMethod_inherit_option is called when entering the method_inherit_option production.
	EnterMethod_inherit_option(c *Method_inherit_optionContext)

	// EnterFinal_inst_list_options is called when entering the final_inst_list_options production.
	EnterFinal_inst_list_options(c *Final_inst_list_optionsContext)

	// EnterFinal_inst_list is called when entering the final_inst_list production.
	EnterFinal_inst_list(c *Final_inst_listContext)

	// EnterFinal_inst is called when entering the final_inst production.
	EnterFinal_inst(c *Final_instContext)

	// EnterOverriding_option is called when entering the overriding_option production.
	EnterOverriding_option(c *Overriding_optionContext)

	// EnterMethod_attr_options is called when entering the method_attr_options production.
	EnterMethod_attr_options(c *Method_attr_optionsContext)

	// EnterMethod_attr is called when entering the method_attr production.
	EnterMethod_attr(c *Method_attrContext)

	// EnterCreate_type_body_stmt is called when entering the create_type_body_stmt production.
	EnterCreate_type_body_stmt(c *Create_type_body_stmtContext)

	// EnterObject_body_def_list is called when entering the object_body_def_list production.
	EnterObject_body_def_list(c *Object_body_def_listContext)

	// EnterObject_body_def is called when entering the object_body_def production.
	EnterObject_body_def(c *Object_body_defContext)

	// EnterCreate_context_stmt is called when entering the create_context_stmt production.
	EnterCreate_context_stmt(c *Create_context_stmtContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterInitialized is called when entering the initialized production.
	EnterInitialized(c *InitializedContext)

	// EnterCreate_directory_stmt is called when entering the create_directory_stmt production.
	EnterCreate_directory_stmt(c *Create_directory_stmtContext)

	// EnterCreate_crypto_stmt is called when entering the create_crypto_stmt production.
	EnterCreate_crypto_stmt(c *Create_crypto_stmtContext)

	// EnterAlter_crypto_stmt is called when entering the alter_crypto_stmt production.
	EnterAlter_crypto_stmt(c *Alter_crypto_stmtContext)

	// EnterAlter_crypto_action is called when entering the alter_crypto_action production.
	EnterAlter_crypto_action(c *Alter_crypto_actionContext)

	// EnterComment_stmt is called when entering the comment_stmt production.
	EnterComment_stmt(c *Comment_stmtContext)

	// EnterComment_tag is called when entering the comment_tag production.
	EnterComment_tag(c *Comment_tagContext)

	// EnterCreate_partition_group_stmt is called when entering the create_partition_group_stmt production.
	EnterCreate_partition_group_stmt(c *Create_partition_group_stmtContext)

	// EnterStorage_act_datatype is called when entering the storage_act_datatype production.
	EnterStorage_act_datatype(c *Storage_act_datatypeContext)

	// EnterPg_storage_lst is called when entering the pg_storage_lst production.
	EnterPg_storage_lst(c *Pg_storage_lstContext)

	// EnterPg_storage is called when entering the pg_storage production.
	EnterPg_storage(c *Pg_storageContext)

	// EnterCreate_table_stmt is called when entering the create_table_stmt production.
	EnterCreate_table_stmt(c *Create_table_stmtContext)

	// EnterCtab_append_attr_clause is called when entering the ctab_append_attr_clause production.
	EnterCtab_append_attr_clause(c *Ctab_append_attr_clauseContext)

	// EnterCtab_append_attr_list is called when entering the ctab_append_attr_list production.
	EnterCtab_append_attr_list(c *Ctab_append_attr_listContext)

	// EnterCobjtab_append_attr_clause is called when entering the cobjtab_append_attr_clause production.
	EnterCobjtab_append_attr_clause(c *Cobjtab_append_attr_clauseContext)

	// EnterCobjtab_append_attr_list is called when entering the cobjtab_append_attr_list production.
	EnterCobjtab_append_attr_list(c *Cobjtab_append_attr_listContext)

	// EnterCtab_append_attr is called when entering the ctab_append_attr production.
	EnterCtab_append_attr(c *Ctab_append_attrContext)

	// EnterCobjtab_append_attr is called when entering the cobjtab_append_attr production.
	EnterCobjtab_append_attr(c *Cobjtab_append_attrContext)

	// EnterCreate_table_action is called when entering the create_table_action production.
	EnterCreate_table_action(c *Create_table_actionContext)

	// EnterCtab_log_options is called when entering the ctab_log_options production.
	EnterCtab_log_options(c *Ctab_log_optionsContext)

	// EnterCtab_log_option is called when entering the ctab_log_option production.
	EnterCtab_log_option(c *Ctab_log_optionContext)

	// EnterCtab_error_options is called when entering the ctab_error_options production.
	EnterCtab_error_options(c *Ctab_error_optionsContext)

	// EnterAdvance_log_clause is called when entering the advance_log_clause production.
	EnterAdvance_log_clause(c *Advance_log_clauseContext)

	// EnterAdd_log_clause is called when entering the add_log_clause production.
	EnterAdd_log_clause(c *Add_log_clauseContext)

	// EnterCtab_error_option is called when entering the ctab_error_option production.
	EnterCtab_error_option(c *Ctab_error_optionContext)

	// EnterCtab_row_movement_clause is called when entering the ctab_row_movement_clause production.
	EnterCtab_row_movement_clause(c *Ctab_row_movement_clauseContext)

	// EnterRange_distribute_act is called when entering the range_distribute_act production.
	EnterRange_distribute_act(c *Range_distribute_actContext)

	// EnterRange_distribute_act_lst is called when entering the range_distribute_act_lst production.
	EnterRange_distribute_act_lst(c *Range_distribute_act_lstContext)

	// EnterList_distribute_act is called when entering the list_distribute_act production.
	EnterList_distribute_act(c *List_distribute_actContext)

	// EnterList_distribute_act_list is called when entering the list_distribute_act_list production.
	EnterList_distribute_act_list(c *List_distribute_act_listContext)

	// EnterDistribute_by_option is called when entering the distribute_by_option production.
	EnterDistribute_by_option(c *Distribute_by_optionContext)

	// EnterDistribute_by is called when entering the distribute_by production.
	EnterDistribute_by(c *Distribute_byContext)

	// EnterIncrement_set is called when entering the increment_set production.
	EnterIncrement_set(c *Increment_setContext)

	// EnterIncrement is called when entering the increment production.
	EnterIncrement(c *IncrementContext)

	// EnterRowdependencies_clause is called when entering the rowdependencies_clause production.
	EnterRowdependencies_clause(c *Rowdependencies_clauseContext)

	// EnterPg_sub_partition is called when entering the pg_sub_partition production.
	EnterPg_sub_partition(c *Pg_sub_partitionContext)

	// EnterTable_type_option is called when entering the table_type_option production.
	EnterTable_type_option(c *Table_type_optionContext)

	// EnterTable_temp_option is called when entering the table_temp_option production.
	EnterTable_temp_option(c *Table_temp_optionContext)

	// EnterObjtab_elem_constraint is called when entering the objtab_elem_constraint production.
	EnterObjtab_elem_constraint(c *Objtab_elem_constraintContext)

	// EnterObjtab_element_cons_list is called when entering the objtab_element_cons_list production.
	EnterObjtab_element_cons_list(c *Objtab_element_cons_listContext)

	// EnterObjtab_element_cons is called when entering the objtab_element_cons production.
	EnterObjtab_element_cons(c *Objtab_element_consContext)

	// EnterObjcol_constraint is called when entering the objcol_constraint production.
	EnterObjcol_constraint(c *Objcol_constraintContext)

	// EnterTable_element_list is called when entering the table_element_list production.
	EnterTable_element_list(c *Table_element_listContext)

	// EnterTable_element is called when entering the table_element production.
	EnterTable_element(c *Table_elementContext)

	// EnterTable_constraint_def is called when entering the table_constraint_def production.
	EnterTable_constraint_def(c *Table_constraint_defContext)

	// EnterOn_commit_option is called when entering the on_commit_option production.
	EnterOn_commit_option(c *On_commit_optionContext)

	// EnterOn_commit_option_nn is called when entering the on_commit_option_nn production.
	EnterOn_commit_option_nn(c *On_commit_option_nnContext)

	// EnterLogging_option is called when entering the logging_option production.
	EnterLogging_option(c *Logging_optionContext)

	// EnterLogging_option_nn is called when entering the logging_option_nn production.
	EnterLogging_option_nn(c *Logging_option_nnContext)

	// EnterPartition_clause is called when entering the partition_clause production.
	EnterPartition_clause(c *Partition_clauseContext)

	// EnterPartition_clause_nn is called when entering the partition_clause_nn production.
	EnterPartition_clause_nn(c *Partition_clause_nnContext)

	// EnterHorizon_partition_clause is called when entering the horizon_partition_clause production.
	EnterHorizon_partition_clause(c *Horizon_partition_clauseContext)

	// EnterCompress_tag_hdr is called when entering the compress_tag_hdr production.
	EnterCompress_tag_hdr(c *Compress_tag_hdrContext)

	// EnterCompress_clause_opt is called when entering the compress_clause_opt production.
	EnterCompress_clause_opt(c *Compress_clause_optContext)

	// EnterCompress_tag is called when entering the compress_tag production.
	EnterCompress_tag(c *Compress_tagContext)

	// EnterCompress_level is called when entering the compress_level production.
	EnterCompress_level(c *Compress_levelContext)

	// EnterCompress_type is called when entering the compress_type production.
	EnterCompress_type(c *Compress_typeContext)

	// EnterRange_partition is called when entering the range_partition production.
	EnterRange_partition(c *Range_partitionContext)

	// EnterRange_partition_list is called when entering the range_partition_list production.
	EnterRange_partition_list(c *Range_partition_listContext)

	// EnterHash_partition is called when entering the hash_partition production.
	EnterHash_partition(c *Hash_partitionContext)

	// EnterHash_partition_list is called when entering the hash_partition_list production.
	EnterHash_partition_list(c *Hash_partition_listContext)

	// EnterList_partition is called when entering the list_partition production.
	EnterList_partition(c *List_partitionContext)

	// EnterList_partition_list is called when entering the list_partition_list production.
	EnterList_partition_list(c *List_partition_listContext)

	// EnterSplit_partition_list is called when entering the split_partition_list production.
	EnterSplit_partition_list(c *Split_partition_listContext)

	// EnterPartition_act is called when entering the partition_act production.
	EnterPartition_act(c *Partition_actContext)

	// EnterVertical_partition_act is called when entering the vertical_partition_act production.
	EnterVertical_partition_act(c *Vertical_partition_actContext)

	// EnterInterval_item is called when entering the interval_item production.
	EnterInterval_item(c *Interval_itemContext)

	// EnterHorizon_partition_act_datatype is called when entering the horizon_partition_act_datatype production.
	EnterHorizon_partition_act_datatype(c *Horizon_partition_act_datatypeContext)

	// EnterHorizon_partition_act is called when entering the horizon_partition_act production.
	EnterHorizon_partition_act(c *Horizon_partition_actContext)

	// EnterLock_partitions_clause is called when entering the lock_partitions_clause production.
	EnterLock_partitions_clause(c *Lock_partitions_clauseContext)

	// EnterSubpartion_template_list_datatype_options is called when entering the subpartion_template_list_datatype_options production.
	EnterSubpartion_template_list_datatype_options(c *Subpartion_template_list_datatype_optionsContext)

	// EnterSubpartion_template_list_datatype is called when entering the subpartion_template_list_datatype production.
	EnterSubpartion_template_list_datatype(c *Subpartion_template_list_datatypeContext)

	// EnterSubpartion_template_list_options is called when entering the subpartion_template_list_options production.
	EnterSubpartion_template_list_options(c *Subpartion_template_list_optionsContext)

	// EnterSubpartion_template_list is called when entering the subpartion_template_list production.
	EnterSubpartion_template_list(c *Subpartion_template_listContext)

	// EnterSubpartion_template_datatype is called when entering the subpartion_template_datatype production.
	EnterSubpartion_template_datatype(c *Subpartion_template_datatypeContext)

	// EnterRange_subpartion_template_datatype is called when entering the range_subpartion_template_datatype production.
	EnterRange_subpartion_template_datatype(c *Range_subpartion_template_datatypeContext)

	// EnterHash_subpartion_template_datatype is called when entering the hash_subpartion_template_datatype production.
	EnterHash_subpartion_template_datatype(c *Hash_subpartion_template_datatypeContext)

	// EnterHash_subpartions_template_datatype_option is called when entering the hash_subpartions_template_datatype_option production.
	EnterHash_subpartions_template_datatype_option(c *Hash_subpartions_template_datatype_optionContext)

	// EnterList_subpartion_template_datatype is called when entering the list_subpartion_template_datatype production.
	EnterList_subpartion_template_datatype(c *List_subpartion_template_datatypeContext)

	// EnterSubpartion_template is called when entering the subpartion_template production.
	EnterSubpartion_template(c *Subpartion_templateContext)

	// EnterRange_subpartion_template is called when entering the range_subpartion_template production.
	EnterRange_subpartion_template(c *Range_subpartion_templateContext)

	// EnterHash_subpartion_template is called when entering the hash_subpartion_template production.
	EnterHash_subpartion_template(c *Hash_subpartion_templateContext)

	// EnterHash_subpartions_template_option is called when entering the hash_subpartions_template_option production.
	EnterHash_subpartions_template_option(c *Hash_subpartions_template_optionContext)

	// EnterList_subpartion_template is called when entering the list_subpartion_template production.
	EnterList_subpartion_template(c *List_subpartion_templateContext)

	// EnterRange_subpartition is called when entering the range_subpartition production.
	EnterRange_subpartition(c *Range_subpartitionContext)

	// EnterHash_subpartition is called when entering the hash_subpartition production.
	EnterHash_subpartition(c *Hash_subpartitionContext)

	// EnterList_subpartition is called when entering the list_subpartition production.
	EnterList_subpartition(c *List_subpartitionContext)

	// EnterRange_subpartition_list is called when entering the range_subpartition_list production.
	EnterRange_subpartition_list(c *Range_subpartition_listContext)

	// EnterHash_subpartition_list is called when entering the hash_subpartition_list production.
	EnterHash_subpartition_list(c *Hash_subpartition_listContext)

	// EnterList_subpartition_list is called when entering the list_subpartition_list production.
	EnterList_subpartition_list(c *List_subpartition_listContext)

	// EnterSubpartition_hash_desc is called when entering the subpartition_hash_desc production.
	EnterSubpartition_hash_desc(c *Subpartition_hash_descContext)

	// EnterSubpartition_range_desc is called when entering the subpartition_range_desc production.
	EnterSubpartition_range_desc(c *Subpartition_range_descContext)

	// EnterSubpartition_list_desc is called when entering the subpartition_list_desc production.
	EnterSubpartition_list_desc(c *Subpartition_list_descContext)

	// EnterSubpartition_hash_desc_list is called when entering the subpartition_hash_desc_list production.
	EnterSubpartition_hash_desc_list(c *Subpartition_hash_desc_listContext)

	// EnterSubpartition_range_desc_list is called when entering the subpartition_range_desc_list production.
	EnterSubpartition_range_desc_list(c *Subpartition_range_desc_listContext)

	// EnterSubpartition_list_desc_list is called when entering the subpartition_list_desc_list production.
	EnterSubpartition_list_desc_list(c *Subpartition_list_desc_listContext)

	// EnterSubpartition_desc is called when entering the subpartition_desc production.
	EnterSubpartition_desc(c *Subpartition_descContext)

	// EnterSubpartition_desc_option is called when entering the subpartition_desc_option production.
	EnterSubpartition_desc_option(c *Subpartition_desc_optionContext)

	// EnterAdd_subpartition_desc is called when entering the add_subpartition_desc production.
	EnterAdd_subpartition_desc(c *Add_subpartition_descContext)

	// EnterPartition_no is called when entering the partition_no production.
	EnterPartition_no(c *Partition_noContext)

	// EnterComment_clause is called when entering the comment_clause production.
	EnterComment_clause(c *Comment_clauseContext)

	// EnterEncrypt_clause_options is called when entering the encrypt_clause_options production.
	EnterEncrypt_clause_options(c *Encrypt_clause_optionsContext)

	// EnterEncrypt_clause is called when entering the encrypt_clause production.
	EnterEncrypt_clause(c *Encrypt_clauseContext)

	// EnterEncrypt_cipher is called when entering the encrypt_cipher production.
	EnterEncrypt_cipher(c *Encrypt_cipherContext)

	// EnterCrypto_name is called when entering the crypto_name production.
	EnterCrypto_name(c *Crypto_nameContext)

	// EnterCipher_name is called when entering the cipher_name production.
	EnterCipher_name(c *Cipher_nameContext)

	// EnterFull_cipher_name is called when entering the full_cipher_name production.
	EnterFull_cipher_name(c *Full_cipher_nameContext)

	// EnterEncrypt_type is called when entering the encrypt_type production.
	EnterEncrypt_type(c *Encrypt_typeContext)

	// EnterManual_clause is called when entering the manual_clause production.
	EnterManual_clause(c *Manual_clauseContext)

	// EnterUser_clause_options is called when entering the user_clause_options production.
	EnterUser_clause_options(c *User_clause_optionsContext)

	// EnterUser_clause is called when entering the user_clause production.
	EnterUser_clause(c *User_clauseContext)

	// EnterUser_list_option is called when entering the user_list_option production.
	EnterUser_list_option(c *User_list_optionContext)

	// EnterUser_list is called when entering the user_list production.
	EnterUser_list(c *User_listContext)

	// EnterHash_cipher is called when entering the hash_cipher production.
	EnterHash_cipher(c *Hash_cipherContext)

	// EnterHash_type is called when entering the hash_type production.
	EnterHash_type(c *Hash_typeContext)

	// EnterSpace_limit is called when entering the space_limit production.
	EnterSpace_limit(c *Space_limitContext)

	// EnterSpace_limit_nn is called when entering the space_limit_nn production.
	EnterSpace_limit_nn(c *Space_limit_nnContext)

	// EnterSpace_limit_1 is called when entering the space_limit_1 production.
	EnterSpace_limit_1(c *Space_limit_1Context)

	// EnterSpace_limit2 is called when entering the space_limit2 production.
	EnterSpace_limit2(c *Space_limit2Context)

	// EnterDel_res is called when entering the del_res production.
	EnterDel_res(c *Del_resContext)

	// EnterTrig_enable is called when entering the trig_enable production.
	EnterTrig_enable(c *Trig_enableContext)

	// EnterAt_raft is called when entering the at_raft production.
	EnterAt_raft(c *At_raftContext)

	// EnterCreate_trigger_stmt is called when entering the create_trigger_stmt production.
	EnterCreate_trigger_stmt(c *Create_trigger_stmtContext)

	// EnterBefore_after is called when entering the before_after production.
	EnterBefore_after(c *Before_afterContext)

	// EnterTrig_del_ins_upd_list is called when entering the trig_del_ins_upd_list production.
	EnterTrig_del_ins_upd_list(c *Trig_del_ins_upd_listContext)

	// EnterTrig_del_ins_upd is called when entering the trig_del_ins_upd production.
	EnterTrig_del_ins_upd(c *Trig_del_ins_updContext)

	// EnterUpdate_of_option is called when entering the update_of_option production.
	EnterUpdate_of_option(c *Update_of_optionContext)

	// EnterNowait is called when entering the nowait production.
	EnterNowait(c *NowaitContext)

	// EnterTrig_event_list is called when entering the trig_event_list production.
	EnterTrig_event_list(c *Trig_event_listContext)

	// EnterTrig_event is called when entering the trig_event production.
	EnterTrig_event(c *Trig_eventContext)

	// EnterEvent_object is called when entering the event_object production.
	EnterEvent_object(c *Event_objectContext)

	// EnterTrig_referencing_def_options is called when entering the trig_referencing_def_options production.
	EnterTrig_referencing_def_options(c *Trig_referencing_def_optionsContext)

	// EnterTrig_referencing_def is called when entering the trig_referencing_def production.
	EnterTrig_referencing_def(c *Trig_referencing_defContext)

	// EnterTrig_referencing_list is called when entering the trig_referencing_list production.
	EnterTrig_referencing_list(c *Trig_referencing_listContext)

	// EnterTrig_referencing_old is called when entering the trig_referencing_old production.
	EnterTrig_referencing_old(c *Trig_referencing_oldContext)

	// EnterTrig_referencing_new is called when entering the trig_referencing_new production.
	EnterTrig_referencing_new(c *Trig_referencing_newContext)

	// EnterTrig_for_each_option is called when entering the trig_for_each_option production.
	EnterTrig_for_each_option(c *Trig_for_each_optionContext)

	// EnterTrig_timer_rate is called when entering the trig_timer_rate production.
	EnterTrig_timer_rate(c *Trig_timer_rateContext)

	// EnterExec_ep_seqno is called when entering the exec_ep_seqno production.
	EnterExec_ep_seqno(c *Exec_ep_seqnoContext)

	// EnterRate_over_day is called when entering the rate_over_day production.
	EnterRate_over_day(c *Rate_over_dayContext)

	// EnterMonth_rate is called when entering the month_rate production.
	EnterMonth_rate(c *Month_rateContext)

	// EnterDay_in_month is called when entering the day_in_month production.
	EnterDay_in_month(c *Day_in_monthContext)

	// EnterDay_in_month_week is called when entering the day_in_month_week production.
	EnterDay_in_month_week(c *Day_in_month_weekContext)

	// EnterWeek_rate is called when entering the week_rate production.
	EnterWeek_rate(c *Week_rateContext)

	// EnterDay_of_week_list is called when entering the day_of_week_list production.
	EnterDay_of_week_list(c *Day_of_week_listContext)

	// EnterDay_rate is called when entering the day_rate production.
	EnterDay_rate(c *Day_rateContext)

	// EnterRate_in_day is called when entering the rate_in_day production.
	EnterRate_in_day(c *Rate_in_dayContext)

	// EnterOnce_in_day is called when entering the once_in_day production.
	EnterOnce_in_day(c *Once_in_dayContext)

	// EnterTimes_in_day is called when entering the times_in_day production.
	EnterTimes_in_day(c *Times_in_dayContext)

	// EnterDuaring_time is called when entering the duaring_time production.
	EnterDuaring_time(c *Duaring_timeContext)

	// EnterDuaring_date is called when entering the duaring_date production.
	EnterDuaring_date(c *Duaring_dateContext)

	// EnterTrig_when_option is called when entering the trig_when_option production.
	EnterTrig_when_option(c *Trig_when_optionContext)

	// EnterTrig_when_condition is called when entering the trig_when_condition production.
	EnterTrig_when_condition(c *Trig_when_conditionContext)

	// EnterRepeat_interval_stmt is called when entering the repeat_interval_stmt production.
	EnterRepeat_interval_stmt(c *Repeat_interval_stmtContext)

	// EnterMax_run_duration is called when entering the max_run_duration production.
	EnterMax_run_duration(c *Max_run_durationContext)

	// EnterRepeat_interval is called when entering the repeat_interval production.
	EnterRepeat_interval(c *Repeat_intervalContext)

	// EnterFrequency_clause is called when entering the frequency_clause production.
	EnterFrequency_clause(c *Frequency_clauseContext)

	// EnterFrequency_define is called when entering the frequency_define production.
	EnterFrequency_define(c *Frequency_defineContext)

	// EnterPredefined_frequency is called when entering the predefined_frequency production.
	EnterPredefined_frequency(c *Predefined_frequencyContext)

	// EnterInterval_clause_list is called when entering the interval_clause_list production.
	EnterInterval_clause_list(c *Interval_clause_listContext)

	// EnterInterval_clause_single is called when entering the interval_clause_single production.
	EnterInterval_clause_single(c *Interval_clause_singleContext)

	// EnterInterval_clause is called when entering the interval_clause production.
	EnterInterval_clause(c *Interval_clauseContext)

	// EnterIntervalnum is called when entering the intervalnum production.
	EnterIntervalnum(c *IntervalnumContext)

	// EnterBymonth_clause is called when entering the bymonth_clause production.
	EnterBymonth_clause(c *Bymonth_clauseContext)

	// EnterMonthlist is called when entering the monthlist production.
	EnterMonthlist(c *MonthlistContext)

	// EnterMonth is called when entering the month production.
	EnterMonth(c *MonthContext)

	// EnterNumeric_month is called when entering the numeric_month production.
	EnterNumeric_month(c *Numeric_monthContext)

	// EnterChar_month is called when entering the char_month production.
	EnterChar_month(c *Char_monthContext)

	// EnterByweekno_clause is called when entering the byweekno_clause production.
	EnterByweekno_clause(c *Byweekno_clauseContext)

	// EnterWeekno_list is called when entering the weekno_list production.
	EnterWeekno_list(c *Weekno_listContext)

	// EnterWeekno is called when entering the weekno production.
	EnterWeekno(c *WeeknoContext)

	// EnterByyearday_clause is called when entering the byyearday_clause production.
	EnterByyearday_clause(c *Byyearday_clauseContext)

	// EnterYearday_list is called when entering the yearday_list production.
	EnterYearday_list(c *Yearday_listContext)

	// EnterYearday is called when entering the yearday production.
	EnterYearday(c *YeardayContext)

	// EnterBymonthday_clause is called when entering the bymonthday_clause production.
	EnterBymonthday_clause(c *Bymonthday_clauseContext)

	// EnterMonthday_list is called when entering the monthday_list production.
	EnterMonthday_list(c *Monthday_listContext)

	// EnterMonthday is called when entering the monthday production.
	EnterMonthday(c *MonthdayContext)

	// EnterByday_clause is called when entering the byday_clause production.
	EnterByday_clause(c *Byday_clauseContext)

	// EnterByday_list is called when entering the byday_list production.
	EnterByday_list(c *Byday_listContext)

	// EnterByday is called when entering the byday production.
	EnterByday(c *BydayContext)

	// EnterWeekdaynum_options is called when entering the weekdaynum_options production.
	EnterWeekdaynum_options(c *Weekdaynum_optionsContext)

	// EnterWeekdaynum is called when entering the weekdaynum production.
	EnterWeekdaynum(c *WeekdaynumContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// EnterByhour_clause is called when entering the byhour_clause production.
	EnterByhour_clause(c *Byhour_clauseContext)

	// EnterHour_list is called when entering the hour_list production.
	EnterHour_list(c *Hour_listContext)

	// EnterHour is called when entering the hour production.
	EnterHour(c *HourContext)

	// EnterByminute_clause is called when entering the byminute_clause production.
	EnterByminute_clause(c *Byminute_clauseContext)

	// EnterMinute_list is called when entering the minute_list production.
	EnterMinute_list(c *Minute_listContext)

	// EnterMinute is called when entering the minute production.
	EnterMinute(c *MinuteContext)

	// EnterBysecond_clause is called when entering the bysecond_clause production.
	EnterBysecond_clause(c *Bysecond_clauseContext)

	// EnterSecond_list is called when entering the second_list production.
	EnterSecond_list(c *Second_listContext)

	// EnterSecond is called when entering the second production.
	EnterSecond(c *SecondContext)

	// EnterQuery_rewrite is called when entering the query_rewrite production.
	EnterQuery_rewrite(c *Query_rewriteContext)

	// EnterBuild_clause is called when entering the build_clause production.
	EnterBuild_clause(c *Build_clauseContext)

	// EnterMv_refresh_option is called when entering the mv_refresh_option production.
	EnterMv_refresh_option(c *Mv_refresh_optionContext)

	// EnterMv_refresh_option_list is called when entering the mv_refresh_option_list production.
	EnterMv_refresh_option_list(c *Mv_refresh_option_listContext)

	// EnterMv_refresh_clause is called when entering the mv_refresh_clause production.
	EnterMv_refresh_clause(c *Mv_refresh_clauseContext)

	// EnterMv_log_purge_syn_asyn_clause is called when entering the mv_log_purge_syn_asyn_clause production.
	EnterMv_log_purge_syn_asyn_clause(c *Mv_log_purge_syn_asyn_clauseContext)

	// EnterMv_log_purge_clause is called when entering the mv_log_purge_clause production.
	EnterMv_log_purge_clause(c *Mv_log_purge_clauseContext)

	// EnterMv_log_with_syntax_item is called when entering the mv_log_with_syntax_item production.
	EnterMv_log_with_syntax_item(c *Mv_log_with_syntax_itemContext)

	// EnterMv_log_with_syntax_item_list is called when entering the mv_log_with_syntax_item_list production.
	EnterMv_log_with_syntax_item_list(c *Mv_log_with_syntax_item_listContext)

	// EnterMv_log_including_new_values is called when entering the mv_log_including_new_values production.
	EnterMv_log_including_new_values(c *Mv_log_including_new_valuesContext)

	// EnterMv_log_with_clause_null is called when entering the mv_log_with_clause_null production.
	EnterMv_log_with_clause_null(c *Mv_log_with_clause_nullContext)

	// EnterCreate_materialized_view_log_stmt is called when entering the create_materialized_view_log_stmt production.
	EnterCreate_materialized_view_log_stmt(c *Create_materialized_view_log_stmtContext)

	// EnterPrebuilt_table_clause_null is called when entering the prebuilt_table_clause_null production.
	EnterPrebuilt_table_clause_null(c *Prebuilt_table_clause_nullContext)

	// EnterCreate_materialized_view_stmt is called when entering the create_materialized_view_stmt production.
	EnterCreate_materialized_view_stmt(c *Create_materialized_view_stmtContext)

	// EnterCreate_view_stmt is called when entering the create_view_stmt production.
	EnterCreate_view_stmt(c *Create_view_stmtContext)

	// EnterCreate_view_stmt_body is called when entering the create_view_stmt_body production.
	EnterCreate_view_stmt_body(c *Create_view_stmt_bodyContext)

	// EnterColumn_list3_options is called when entering the column_list3_options production.
	EnterColumn_list3_options(c *Column_list3_optionsContext)

	// EnterColumn_list3 is called when entering the column_list3 production.
	EnterColumn_list3(c *Column_list3Context)

	// EnterView_column_constraint_def is called when entering the view_column_constraint_def production.
	EnterView_column_constraint_def(c *View_column_constraint_defContext)

	// EnterView_column_constraints is called when entering the view_column_constraints production.
	EnterView_column_constraints(c *View_column_constraintsContext)

	// EnterView_column_constraint is called when entering the view_column_constraint production.
	EnterView_column_constraint(c *View_column_constraintContext)

	// EnterView_column_constraint_action is called when entering the view_column_constraint_action production.
	EnterView_column_constraint_action(c *View_column_constraint_actionContext)

	// EnterView_constraint_def is called when entering the view_constraint_def production.
	EnterView_constraint_def(c *View_constraint_defContext)

	// EnterWith_schemabinding is called when entering the with_schemabinding production.
	EnterWith_schemabinding(c *With_schemabindingContext)

	// EnterColumn_list_options is called when entering the column_list_options production.
	EnterColumn_list_options(c *Column_list_optionsContext)

	// EnterWith_check_or_readonly_option is called when entering the with_check_or_readonly_option production.
	EnterWith_check_or_readonly_option(c *With_check_or_readonly_optionContext)

	// EnterCheck_level_option is called when entering the check_level_option production.
	EnterCheck_level_option(c *Check_level_optionContext)

	// EnterDecl_cursor is called when entering the decl_cursor production.
	EnterDecl_cursor(c *Decl_cursorContext)

	// EnterDelete_stmt is called when entering the delete_stmt production.
	EnterDelete_stmt(c *Delete_stmtContext)

	// EnterDelete_stmt_body is called when entering the delete_stmt_body production.
	EnterDelete_stmt_body(c *Delete_stmt_bodyContext)

	// EnterDelete_multi_tv_option is called when entering the delete_multi_tv_option production.
	EnterDelete_multi_tv_option(c *Delete_multi_tv_optionContext)

	// EnterCheck_limit_option is called when entering the check_limit_option production.
	EnterCheck_limit_option(c *Check_limit_optionContext)

	// EnterWhere_current_option is called when entering the where_current_option production.
	EnterWhere_current_option(c *Where_current_optionContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterStart_with_clause_null is called when entering the start_with_clause_null production.
	EnterStart_with_clause_null(c *Start_with_clause_nullContext)

	// EnterConnect_by_item is called when entering the connect_by_item production.
	EnterConnect_by_item(c *Connect_by_itemContext)

	// EnterConnect_by_clause is called when entering the connect_by_clause production.
	EnterConnect_by_clause(c *Connect_by_clauseContext)

	// EnterHierarchical_query_clause is called when entering the hierarchical_query_clause production.
	EnterHierarchical_query_clause(c *Hierarchical_query_clauseContext)

	// EnterNocycle_flag is called when entering the nocycle_flag production.
	EnterNocycle_flag(c *Nocycle_flagContext)

	// EnterSearch_condition is called when entering the search_condition production.
	EnterSearch_condition(c *Search_conditionContext)

	// EnterDisconnect_stmt is called when entering the disconnect_stmt production.
	EnterDisconnect_stmt(c *Disconnect_stmtContext)

	// EnterDisconnect_option is called when entering the disconnect_option production.
	EnterDisconnect_option(c *Disconnect_optionContext)

	// EnterDrop_stmt is called when entering the drop_stmt production.
	EnterDrop_stmt(c *Drop_stmtContext)

	// EnterDrop_stmt_body_1 is called when entering the drop_stmt_body_1 production.
	EnterDrop_stmt_body_1(c *Drop_stmt_body_1Context)

	// EnterDrop_stmt_2 is called when entering the drop_stmt_2 production.
	EnterDrop_stmt_2(c *Drop_stmt_2Context)

	// EnterDrop_id_db_object is called when entering the drop_id_db_object production.
	EnterDrop_id_db_object(c *Drop_id_db_objectContext)

	// EnterId_db_object is called when entering the id_db_object production.
	EnterId_db_object(c *Id_db_objectContext)

	// EnterDrop_db_object is called when entering the drop_db_object production.
	EnterDrop_db_object(c *Drop_db_objectContext)

	// EnterExist is called when entering the exist production.
	EnterExist(c *ExistContext)

	// EnterNot_exist is called when entering the not_exist production.
	EnterNot_exist(c *Not_existContext)

	// EnterDb_object is called when entering the db_object production.
	EnterDb_object(c *Db_objectContext)

	// EnterIs_detach is called when entering the is_detach production.
	EnterIs_detach(c *Is_detachContext)

	// EnterPurge_option is called when entering the purge_option production.
	EnterPurge_option(c *Purge_optionContext)

	// EnterAlter_database_stmt is called when entering the alter_database_stmt production.
	EnterAlter_database_stmt(c *Alter_database_stmtContext)

	// EnterAlter_system_action is called when entering the alter_system_action production.
	EnterAlter_system_action(c *Alter_system_actionContext)

	// EnterAlter_database_action is called when entering the alter_database_action production.
	EnterAlter_database_action(c *Alter_database_actionContext)

	// EnterForce is called when entering the force production.
	EnterForce(c *ForceContext)

	// EnterTablespace_name is called when entering the tablespace_name production.
	EnterTablespace_name(c *Tablespace_nameContext)

	// EnterRaft_name is called when entering the raft_name production.
	EnterRaft_name(c *Raft_nameContext)

	// EnterFetch_into is called when entering the fetch_into production.
	EnterFetch_into(c *Fetch_intoContext)

	// EnterBulk_or_single_into is called when entering the bulk_or_single_into production.
	EnterBulk_or_single_into(c *Bulk_or_single_intoContext)

	// EnterFetch_stmt is called when entering the fetch_stmt production.
	EnterFetch_stmt(c *Fetch_stmtContext)

	// EnterFetch_statement is called when entering the fetch_statement production.
	EnterFetch_statement(c *Fetch_statementContext)

	// EnterFetch_tail is called when entering the fetch_tail production.
	EnterFetch_tail(c *Fetch_tailContext)

	// EnterFetch_limit_option is called when entering the fetch_limit_option production.
	EnterFetch_limit_option(c *Fetch_limit_optionContext)

	// EnterFetch_option is called when entering the fetch_option production.
	EnterFetch_option(c *Fetch_optionContext)

	// EnterFetch_direct_option is called when entering the fetch_direct_option production.
	EnterFetch_direct_option(c *Fetch_direct_optionContext)

	// EnterLog_errors_into is called when entering the log_errors_into production.
	EnterLog_errors_into(c *Log_errors_intoContext)

	// EnterLog_errors_expression is called when entering the log_errors_expression production.
	EnterLog_errors_expression(c *Log_errors_expressionContext)

	// EnterLog_errors_unlimited is called when entering the log_errors_unlimited production.
	EnterLog_errors_unlimited(c *Log_errors_unlimitedContext)

	// EnterLog_errors is called when entering the log_errors production.
	EnterLog_errors(c *Log_errorsContext)

	// EnterInsert_stmt is called when entering the insert_stmt production.
	EnterInsert_stmt(c *Insert_stmtContext)

	// EnterInsert_stmt_body is called when entering the insert_stmt_body production.
	EnterInsert_stmt_body(c *Insert_stmt_bodyContext)

	// EnterFull_column_list_options is called when entering the full_column_list_options production.
	EnterFull_column_list_options(c *Full_column_list_optionsContext)

	// EnterIns_value_options is called when entering the ins_value_options production.
	EnterIns_value_options(c *Ins_value_optionsContext)

	// EnterInsert_into_single is called when entering the insert_into_single production.
	EnterInsert_into_single(c *Insert_into_singleContext)

	// EnterMulti_insert_into_list is called when entering the multi_insert_into_list production.
	EnterMulti_insert_into_list(c *Multi_insert_into_listContext)

	// EnterMulti_insert_tag is called when entering the multi_insert_tag production.
	EnterMulti_insert_tag(c *Multi_insert_tagContext)

	// EnterInsert_into_single_condition is called when entering the insert_into_single_condition production.
	EnterInsert_into_single_condition(c *Insert_into_single_conditionContext)

	// EnterMulti_insert_into_condition_list is called when entering the multi_insert_into_condition_list production.
	EnterMulti_insert_into_condition_list(c *Multi_insert_into_condition_listContext)

	// EnterMulti_insert_into_else is called when entering the multi_insert_into_else production.
	EnterMulti_insert_into_else(c *Multi_insert_into_elseContext)

	// EnterMulti_insert_stmt_body is called when entering the multi_insert_stmt_body production.
	EnterMulti_insert_stmt_body(c *Multi_insert_stmt_bodyContext)

	// EnterInsert_tail is called when entering the insert_tail production.
	EnterInsert_tail(c *Insert_tailContext)

	// EnterInsert_action is called when entering the insert_action production.
	EnterInsert_action(c *Insert_actionContext)

	// EnterRecord_var_values is called when entering the record_var_values production.
	EnterRecord_var_values(c *Record_var_valuesContext)

	// EnterRecord_var_value is called when entering the record_var_value production.
	EnterRecord_var_value(c *Record_var_valueContext)

	// EnterIns_value is called when entering the ins_value production.
	EnterIns_value(c *Ins_valueContext)

	// EnterOpen_stmt is called when entering the open_stmt production.
	EnterOpen_stmt(c *Open_stmtContext)

	// EnterOpen_statement is called when entering the open_statement production.
	EnterOpen_statement(c *Open_statementContext)

	// EnterOpen_tail is called when entering the open_tail production.
	EnterOpen_tail(c *Open_tailContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterRaise_stmt is called when entering the raise_stmt production.
	EnterRaise_stmt(c *Raise_stmtContext)

	// EnterRollback_stmt is called when entering the rollback_stmt production.
	EnterRollback_stmt(c *Rollback_stmtContext)

	// EnterTo_savepoint is called when entering the to_savepoint production.
	EnterTo_savepoint(c *To_savepointContext)

	// EnterSavepoint_stmt is called when entering the savepoint_stmt production.
	EnterSavepoint_stmt(c *Savepoint_stmtContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterAll_distinct_option is called when entering the all_distinct_option production.
	EnterAll_distinct_option(c *All_distinct_optionContext)

	// EnterAll_distinct_option_2 is called when entering the all_distinct_option_2 production.
	EnterAll_distinct_option_2(c *All_distinct_option_2Context)

	// EnterCorresponding_clause is called when entering the corresponding_clause production.
	EnterCorresponding_clause(c *Corresponding_clauseContext)

	// EnterTop_option is called when entering the top_option production.
	EnterTop_option(c *Top_optionContext)

	// EnterLimit_option is called when entering the limit_option production.
	EnterLimit_option(c *Limit_optionContext)

	// EnterLimit_clause is called when entering the limit_clause production.
	EnterLimit_clause(c *Limit_clauseContext)

	// EnterLimit_not_empty is called when entering the limit_not_empty production.
	EnterLimit_not_empty(c *Limit_not_emptyContext)

	// EnterRow_limiting_clause is called when entering the row_limiting_clause production.
	EnterRow_limiting_clause(c *Row_limiting_clauseContext)

	// EnterRow_num_desc is called when entering the row_num_desc production.
	EnterRow_num_desc(c *Row_num_descContext)

	// EnterFirst_next_desc is called when entering the first_next_desc production.
	EnterFirst_next_desc(c *First_next_descContext)

	// EnterSelect_item_list is called when entering the select_item_list production.
	EnterSelect_item_list(c *Select_item_listContext)

	// EnterSelect_item is called when entering the select_item production.
	EnterSelect_item(c *Select_itemContext)

	// EnterAs_alias is called when entering the as_alias production.
	EnterAs_alias(c *As_aliasContext)

	// EnterSelect_tail is called when entering the select_tail production.
	EnterSelect_tail(c *Select_tailContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterFrom_tv_list is called when entering the from_tv_list production.
	EnterFrom_tv_list(c *From_tv_listContext)

	// EnterFrom_tv is called when entering the from_tv production.
	EnterFrom_tv(c *From_tvContext)

	// EnterJoined_table is called when entering the joined_table production.
	EnterJoined_table(c *Joined_tableContext)

	// EnterTrxid is called when entering the trxid production.
	EnterTrxid(c *TrxidContext)

	// EnterFlashback_query_low is called when entering the flashback_query_low production.
	EnterFlashback_query_low(c *Flashback_query_lowContext)

	// EnterTrxid_option is called when entering the trxid_option production.
	EnterTrxid_option(c *Trxid_optionContext)

	// EnterRange_from_to is called when entering the range_from_to production.
	EnterRange_from_to(c *Range_from_toContext)

	// EnterSample_exp is called when entering the sample_exp production.
	EnterSample_exp(c *Sample_expContext)

	// EnterPivot_sfun is called when entering the pivot_sfun production.
	EnterPivot_sfun(c *Pivot_sfunContext)

	// EnterPivot_sfun_lst is called when entering the pivot_sfun_lst production.
	EnterPivot_sfun_lst(c *Pivot_sfun_lstContext)

	// EnterPivot_for_clause is called when entering the pivot_for_clause production.
	EnterPivot_for_clause(c *Pivot_for_clauseContext)

	// EnterPivot_in_clause1_expr is called when entering the pivot_in_clause1_expr production.
	EnterPivot_in_clause1_expr(c *Pivot_in_clause1_exprContext)

	// EnterPivot_in_clause_low_1 is called when entering the pivot_in_clause_low_1 production.
	EnterPivot_in_clause_low_1(c *Pivot_in_clause_low_1Context)

	// EnterPivot_in_clause_low_2 is called when entering the pivot_in_clause_low_2 production.
	EnterPivot_in_clause_low_2(c *Pivot_in_clause_low_2Context)

	// EnterPivot_in_clause_low is called when entering the pivot_in_clause_low production.
	EnterPivot_in_clause_low(c *Pivot_in_clause_lowContext)

	// EnterPivot_xml is called when entering the pivot_xml production.
	EnterPivot_xml(c *Pivot_xmlContext)

	// EnterPivot_clause_low is called when entering the pivot_clause_low production.
	EnterPivot_clause_low(c *Pivot_clause_lowContext)

	// EnterUnpivot_val_col_lst is called when entering the unpivot_val_col_lst production.
	EnterUnpivot_val_col_lst(c *Unpivot_val_col_lstContext)

	// EnterInclude_clause is called when entering the include_clause production.
	EnterInclude_clause(c *Include_clauseContext)

	// EnterUnpivot_in_clause_expr is called when entering the unpivot_in_clause_expr production.
	EnterUnpivot_in_clause_expr(c *Unpivot_in_clause_exprContext)

	// EnterUnpivot_in_clause_low is called when entering the unpivot_in_clause_low production.
	EnterUnpivot_in_clause_low(c *Unpivot_in_clause_lowContext)

	// EnterUnpivot_clause_low is called when entering the unpivot_clause_low production.
	EnterUnpivot_clause_low(c *Unpivot_clause_lowContext)

	// EnterSample_clause_low is called when entering the sample_clause_low production.
	EnterSample_clause_low(c *Sample_clause_lowContext)

	// EnterNormal_tv_name is called when entering the normal_tv_name production.
	EnterNormal_tv_name(c *Normal_tv_nameContext)

	// EnterNormal_tv_tail is called when entering the normal_tv_tail production.
	EnterNormal_tv_tail(c *Normal_tv_tailContext)

	// EnterNormal_tv_tail_low is called when entering the normal_tv_tail_low production.
	EnterNormal_tv_tail_low(c *Normal_tv_tail_lowContext)

	// EnterNormal_alias is called when entering the normal_alias production.
	EnterNormal_alias(c *Normal_aliasContext)

	// EnterNormal_tv_tail_low2 is called when entering the normal_tv_tail_low2 production.
	EnterNormal_tv_tail_low2(c *Normal_tv_tail_low2Context)

	// EnterNormal_tv_tail_low3 is called when entering the normal_tv_tail_low3 production.
	EnterNormal_tv_tail_low3(c *Normal_tv_tail_low3Context)

	// EnterNormal_tv_derived_table_options is called when entering the normal_tv_derived_table_options production.
	EnterNormal_tv_derived_table_options(c *Normal_tv_derived_table_optionsContext)

	// EnterNormal_tv_derived_table_low is called when entering the normal_tv_derived_table_low production.
	EnterNormal_tv_derived_table_low(c *Normal_tv_derived_table_lowContext)

	// EnterNormal_tv_derived_table is called when entering the normal_tv_derived_table production.
	EnterNormal_tv_derived_table(c *Normal_tv_derived_tableContext)

	// EnterSelect_with_paran_with_alias is called when entering the select_with_paran_with_alias production.
	EnterSelect_with_paran_with_alias(c *Select_with_paran_with_aliasContext)

	// EnterFrom_table_exp is called when entering the from_table_exp production.
	EnterFrom_table_exp(c *From_table_expContext)

	// EnterFrom_table_select_with_paran is called when entering the from_table_select_with_paran production.
	EnterFrom_table_select_with_paran(c *From_table_select_with_paranContext)

	// EnterNormal_tv is called when entering the normal_tv production.
	EnterNormal_tv(c *Normal_tvContext)

	// EnterXml_passing is called when entering the xml_passing production.
	EnterXml_passing(c *Xml_passingContext)

	// EnterXmlcoldef_lst_options is called when entering the xmlcoldef_lst_options production.
	EnterXmlcoldef_lst_options(c *Xmlcoldef_lst_optionsContext)

	// EnterXmlcoldef_lst is called when entering the xmlcoldef_lst production.
	EnterXmlcoldef_lst(c *Xmlcoldef_lstContext)

	// EnterXmlcoldef is called when entering the xmlcoldef production.
	EnterXmlcoldef(c *XmlcoldefContext)

	// EnterOn_error is called when entering the on_error production.
	EnterOn_error(c *On_errorContext)

	// EnterJsoncol_lst is called when entering the jsoncol_lst production.
	EnterJsoncol_lst(c *Jsoncol_lstContext)

	// EnterJsoncoldef_lst is called when entering the jsoncoldef_lst production.
	EnterJsoncoldef_lst(c *Jsoncoldef_lstContext)

	// EnterJsoncoldef is called when entering the jsoncoldef production.
	EnterJsoncoldef(c *JsoncoldefContext)

	// EnterJson_exists_col is called when entering the json_exists_col production.
	EnterJson_exists_col(c *Json_exists_colContext)

	// EnterJson_qurey_col is called when entering the json_qurey_col production.
	EnterJson_qurey_col(c *Json_qurey_colContext)

	// EnterJson_value_col is called when entering the json_value_col production.
	EnterJson_value_col(c *Json_value_colContext)

	// EnterJson_nested_col is called when entering the json_nested_col production.
	EnterJson_nested_col(c *Json_nested_colContext)

	// EnterOrdinality_col is called when entering the ordinality_col production.
	EnterOrdinality_col(c *Ordinality_colContext)

	// EnterJoined_table_element is called when entering the joined_table_element production.
	EnterJoined_table_element(c *Joined_table_elementContext)

	// EnterCross_outer_apply_clause is called when entering the cross_outer_apply_clause production.
	EnterCross_outer_apply_clause(c *Cross_outer_apply_clauseContext)

	// EnterCross_outer_apply_join is called when entering the cross_outer_apply_join production.
	EnterCross_outer_apply_join(c *Cross_outer_apply_joinContext)

	// EnterCross_join is called when entering the cross_join production.
	EnterCross_join(c *Cross_joinContext)

	// EnterPartition_out_join is called when entering the partition_out_join production.
	EnterPartition_out_join(c *Partition_out_joinContext)

	// EnterQualified_join is called when entering the qualified_join production.
	EnterQualified_join(c *Qualified_joinContext)

	// EnterQualified_joinspec is called when entering the qualified_joinspec production.
	EnterQualified_joinspec(c *Qualified_joinspecContext)

	// EnterNamed_columns_join is called when entering the named_columns_join production.
	EnterNamed_columns_join(c *Named_columns_joinContext)

	// EnterJoin_hint is called when entering the join_hint production.
	EnterJoin_hint(c *Join_hintContext)

	// EnterJoin_type is called when entering the join_type production.
	EnterJoin_type(c *Join_typeContext)

	// EnterOuter_join_type is called when entering the outer_join_type production.
	EnterOuter_join_type(c *Outer_join_typeContext)

	// EnterJoin_condition is called when entering the join_condition production.
	EnterJoin_condition(c *Join_conditionContext)

	// EnterGroup_by_clause is called when entering the group_by_clause production.
	EnterGroup_by_clause(c *Group_by_clauseContext)

	// EnterGroup_item is called when entering the group_item production.
	EnterGroup_item(c *Group_itemContext)

	// EnterExp_rollup_cube_item2 is called when entering the exp_rollup_cube_item2 production.
	EnterExp_rollup_cube_item2(c *Exp_rollup_cube_item2Context)

	// EnterExp_rollup_cube_item is called when entering the exp_rollup_cube_item production.
	EnterExp_rollup_cube_item(c *Exp_rollup_cube_itemContext)

	// EnterGrouping_set_items is called when entering the grouping_set_items production.
	EnterGrouping_set_items(c *Grouping_set_itemsContext)

	// EnterGrouping_set_item is called when entering the grouping_set_item production.
	EnterGrouping_set_item(c *Grouping_set_itemContext)

	// EnterHaving_clause is called when entering the having_clause production.
	EnterHaving_clause(c *Having_clauseContext)

	// EnterWithout_into_select is called when entering the without_into_select production.
	EnterWithout_into_select(c *Without_into_selectContext)

	// EnterSel_clause_app is called when entering the sel_clause_app production.
	EnterSel_clause_app(c *Sel_clause_appContext)

	// EnterSelect_clause is called when entering the select_clause production.
	EnterSelect_clause(c *Select_clauseContext)

	// EnterSimple_select is called when entering the simple_select production.
	EnterSimple_select(c *Simple_selectContext)

	// EnterSelect_with_paran is called when entering the select_with_paran production.
	EnterSelect_with_paran(c *Select_with_paranContext)

	// EnterQuery_exp is called when entering the query_exp production.
	EnterQuery_exp(c *Query_expContext)

	// EnterFor_xml_path is called when entering the for_xml_path production.
	EnterFor_xml_path(c *For_xml_pathContext)

	// EnterWith_tag is called when entering the with_tag production.
	EnterWith_tag(c *With_tagContext)

	// EnterWith_option is called when entering the with_option production.
	EnterWith_option(c *With_optionContext)

	// EnterWith_clause is called when entering the with_clause production.
	EnterWith_clause(c *With_clauseContext)

	// EnterWith_function_list is called when entering the with_function_list production.
	EnterWith_function_list(c *With_function_listContext)

	// EnterFunc_def_inner is called when entering the func_def_inner production.
	EnterFunc_def_inner(c *Func_def_innerContext)

	// EnterWith_function_list_element is called when entering the with_function_list_element production.
	EnterWith_function_list_element(c *With_function_list_elementContext)

	// EnterWith_view_list is called when entering the with_view_list production.
	EnterWith_view_list(c *With_view_listContext)

	// EnterDepth_type_option is called when entering the depth_type_option production.
	EnterDepth_type_option(c *Depth_type_optionContext)

	// EnterSearch_clause is called when entering the search_clause production.
	EnterSearch_clause(c *Search_clauseContext)

	// EnterCycle_clause is called when entering the cycle_clause production.
	EnterCycle_clause(c *Cycle_clauseContext)

	// EnterWith_view_list_element is called when entering the with_view_list_element production.
	EnterWith_view_list_element(c *With_view_list_elementContext)

	// EnterAssignment_obj_list is called when entering the assignment_obj_list production.
	EnterAssignment_obj_list(c *Assignment_obj_listContext)

	// EnterAssignment_obj is called when entering the assignment_obj production.
	EnterAssignment_obj(c *Assignment_objContext)

	// EnterOrder_by_options is called when entering the order_by_options production.
	EnterOrder_by_options(c *Order_by_optionsContext)

	// EnterOrder_by is called when entering the order_by production.
	EnterOrder_by(c *Order_byContext)

	// EnterAsc_desc_option is called when entering the asc_desc_option production.
	EnterAsc_desc_option(c *Asc_desc_optionContext)

	// EnterNulls_last_option is called when entering the nulls_last_option production.
	EnterNulls_last_option(c *Nulls_last_optionContext)

	// EnterCollate_option is called when entering the collate_option production.
	EnterCollate_option(c *Collate_optionContext)

	// EnterOrder_by_list is called when entering the order_by_list production.
	EnterOrder_by_list(c *Order_by_listContext)

	// EnterOrder_by_item is called when entering the order_by_item production.
	EnterOrder_by_item(c *Order_by_itemContext)

	// EnterFor_update_options is called when entering the for_update_options production.
	EnterFor_update_options(c *For_update_optionsContext)

	// EnterFor_update is called when entering the for_update production.
	EnterFor_update(c *For_updateContext)

	// EnterSet_session_stmt is called when entering the set_session_stmt production.
	EnterSet_session_stmt(c *Set_session_stmtContext)

	// EnterSet_trans_stmt is called when entering the set_trans_stmt production.
	EnterSet_trans_stmt(c *Set_trans_stmtContext)

	// EnterTrans_mode_lstl is called when entering the trans_mode_lstl production.
	EnterTrans_mode_lstl(c *Trans_mode_lstlContext)

	// EnterTrans_mode_lst is called when entering the trans_mode_lst production.
	EnterTrans_mode_lst(c *Trans_mode_lstContext)

	// EnterTrans_mode is called when entering the trans_mode production.
	EnterTrans_mode(c *Trans_modeContext)

	// EnterTime_zone_exp_new is called when entering the time_zone_exp_new production.
	EnterTime_zone_exp_new(c *Time_zone_exp_newContext)

	// EnterTrans_rw_option is called when entering the trans_rw_option production.
	EnterTrans_rw_option(c *Trans_rw_optionContext)

	// EnterTrans_level_option is called when entering the trans_level_option production.
	EnterTrans_level_option(c *Trans_level_optionContext)

	// EnterLock_table_stmt is called when entering the lock_table_stmt production.
	EnterLock_table_stmt(c *Lock_table_stmtContext)

	// EnterLock_mode_option is called when entering the lock_mode_option production.
	EnterLock_mode_option(c *Lock_mode_optionContext)

	// EnterSet_identins_stmt is called when entering the set_identins_stmt production.
	EnterSet_identins_stmt(c *Set_identins_stmtContext)

	// EnterSet_identins_option is called when entering the set_identins_option production.
	EnterSet_identins_option(c *Set_identins_optionContext)

	// EnterTrunc_table_stmt is called when entering the trunc_table_stmt production.
	EnterTrunc_table_stmt(c *Trunc_table_stmtContext)

	// EnterUpdate_stmt is called when entering the update_stmt production.
	EnterUpdate_stmt(c *Update_stmtContext)

	// EnterUpdate_stmt_body is called when entering the update_stmt_body production.
	EnterUpdate_stmt_body(c *Update_stmt_bodyContext)

	// EnterUpdate_tv_list is called when entering the update_tv_list production.
	EnterUpdate_tv_list(c *Update_tv_listContext)

	// EnterReturn_item is called when entering the return_item production.
	EnterReturn_item(c *Return_itemContext)

	// EnterReturn_item_list is called when entering the return_item_list production.
	EnterReturn_item_list(c *Return_item_listContext)

	// EnterReturn_option is called when entering the return_option production.
	EnterReturn_option(c *Return_optionContext)

	// EnterReturn_into_obj is called when entering the return_into_obj production.
	EnterReturn_into_obj(c *Return_into_objContext)

	// EnterCollect_into_rset is called when entering the collect_into_rset production.
	EnterCollect_into_rset(c *Collect_into_rsetContext)

	// EnterAlias_option is called when entering the alias_option production.
	EnterAlias_option(c *Alias_optionContext)

	// EnterSet_value_list is called when entering the set_value_list production.
	EnterSet_value_list(c *Set_value_listContext)

	// EnterSet_value_node is called when entering the set_value_node production.
	EnterSet_value_node(c *Set_value_nodeContext)

	// EnterSet_col_list is called when entering the set_col_list production.
	EnterSet_col_list(c *Set_col_listContext)

	// EnterUpdate_from_clause is called when entering the update_from_clause production.
	EnterUpdate_from_clause(c *Update_from_clauseContext)

	// EnterMerge_into_stmt is called when entering the merge_into_stmt production.
	EnterMerge_into_stmt(c *Merge_into_stmtContext)

	// EnterMerge_into_stmt_body is called when entering the merge_into_stmt_body production.
	EnterMerge_into_stmt_body(c *Merge_into_stmt_bodyContext)

	// EnterMerge_into_sub_clause is called when entering the merge_into_sub_clause production.
	EnterMerge_into_sub_clause(c *Merge_into_sub_clauseContext)

	// EnterMerge_update_clause is called when entering the merge_update_clause production.
	EnterMerge_update_clause(c *Merge_update_clauseContext)

	// EnterMerge_insert_clause is called when entering the merge_insert_clause production.
	EnterMerge_insert_clause(c *Merge_insert_clauseContext)

	// EnterCreate_profile_stmt is called when entering the create_profile_stmt production.
	EnterCreate_profile_stmt(c *Create_profile_stmtContext)

	// EnterAlter_profile_stmt is called when entering the alter_profile_stmt production.
	EnterAlter_profile_stmt(c *Alter_profile_stmtContext)

	// EnterCreate_user_stmt is called when entering the create_user_stmt production.
	EnterCreate_user_stmt(c *Create_user_stmtContext)

	// EnterDefault_ts_name is called when entering the default_ts_name production.
	EnterDefault_ts_name(c *Default_ts_nameContext)

	// EnterDefault_ts_name_lst is called when entering the default_ts_name_lst production.
	EnterDefault_ts_name_lst(c *Default_ts_name_lstContext)

	// EnterDefault_ts_name_node is called when entering the default_ts_name_node production.
	EnterDefault_ts_name_node(c *Default_ts_name_nodeContext)

	// EnterDefault_idx_ts_name is called when entering the default_idx_ts_name production.
	EnterDefault_idx_ts_name(c *Default_idx_ts_nameContext)

	// EnterDefault_ts_name_low is called when entering the default_ts_name_low production.
	EnterDefault_ts_name_low(c *Default_ts_name_lowContext)

	// EnterTemp_ts_name is called when entering the temp_ts_name production.
	EnterTemp_ts_name(c *Temp_ts_nameContext)

	// EnterDefault_ts_group_name_low is called when entering the default_ts_group_name_low production.
	EnterDefault_ts_group_name_low(c *Default_ts_group_name_lowContext)

	// EnterOn_schema is called when entering the on_schema production.
	EnterOn_schema(c *On_schemaContext)

	// EnterReplace_old_pwd is called when entering the replace_old_pwd production.
	EnterReplace_old_pwd(c *Replace_old_pwdContext)

	// EnterAlter_user_stmt is called when entering the alter_user_stmt production.
	EnterAlter_user_stmt(c *Alter_user_stmtContext)

	// EnterUser_encrypt_options is called when entering the user_encrypt_options production.
	EnterUser_encrypt_options(c *User_encrypt_optionsContext)

	// EnterUser_encrypt_option is called when entering the user_encrypt_option production.
	EnterUser_encrypt_option(c *User_encrypt_optionContext)

	// EnterAuthent_type_options is called when entering the authent_type_options production.
	EnterAuthent_type_options(c *Authent_type_optionsContext)

	// EnterHash_cipher_option is called when entering the hash_cipher_option production.
	EnterHash_cipher_option(c *Hash_cipher_optionContext)

	// EnterAuthent_type is called when entering the authent_type production.
	EnterAuthent_type(c *Authent_typeContext)

	// EnterForce_format is called when entering the force_format production.
	EnterForce_format(c *Force_formatContext)

	// EnterAs is called when entering the as production.
	EnterAs(c *AsContext)

	// EnterPwd_policy is called when entering the pwd_policy production.
	EnterPwd_policy(c *Pwd_policyContext)

	// EnterAccount_lock is called when entering the account_lock production.
	EnterAccount_lock(c *Account_lockContext)

	// EnterRead_only_flag is called when entering the read_only_flag production.
	EnterRead_only_flag(c *Read_only_flagContext)

	// EnterRead_only_flag_not_null is called when entering the read_only_flag_not_null production.
	EnterRead_only_flag_not_null(c *Read_only_flag_not_nullContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterExpire is called when entering the expire production.
	EnterExpire(c *ExpireContext)

	// EnterResource_limit_options is called when entering the resource_limit_options production.
	EnterResource_limit_options(c *Resource_limit_optionsContext)

	// EnterResource_limit_list is called when entering the resource_limit_list production.
	EnterResource_limit_list(c *Resource_limit_listContext)

	// EnterResource_limit_list_with_comma is called when entering the resource_limit_list_with_comma production.
	EnterResource_limit_list_with_comma(c *Resource_limit_list_with_commaContext)

	// EnterResource_limit_list_with_empty is called when entering the resource_limit_list_with_empty production.
	EnterResource_limit_list_with_empty(c *Resource_limit_list_with_emptyContext)

	// EnterResource_limit is called when entering the resource_limit production.
	EnterResource_limit(c *Resource_limitContext)

	// EnterResource_limit_value is called when entering the resource_limit_value production.
	EnterResource_limit_value(c *Resource_limit_valueContext)

	// EnterCreate_audit_rule_stmt is called when entering the create_audit_rule_stmt production.
	EnterCreate_audit_rule_stmt(c *Create_audit_rule_stmtContext)

	// EnterRule_name is called when entering the rule_name production.
	EnterRule_name(c *Rule_nameContext)

	// EnterAudit_rule_action is called when entering the audit_rule_action production.
	EnterAudit_rule_action(c *Audit_rule_actionContext)

	// EnterBy_login_or_all is called when entering the by_login_or_all production.
	EnterBy_login_or_all(c *By_login_or_allContext)

	// EnterAllow_ip_list is called when entering the allow_ip_list production.
	EnterAllow_ip_list(c *Allow_ip_listContext)

	// EnterNot_allow_ip_list is called when entering the not_allow_ip_list production.
	EnterNot_allow_ip_list(c *Not_allow_ip_listContext)

	// EnterIp_list is called when entering the ip_list production.
	EnterIp_list(c *Ip_listContext)

	// EnterAllow_dt_list is called when entering the allow_dt_list production.
	EnterAllow_dt_list(c *Allow_dt_listContext)

	// EnterNot_allow_dt_list is called when entering the not_allow_dt_list production.
	EnterNot_allow_dt_list(c *Not_allow_dt_listContext)

	// EnterDt_list is called when entering the dt_list production.
	EnterDt_list(c *Dt_listContext)

	// EnterDt is called when entering the dt production.
	EnterDt(c *DtContext)

	// EnterOp_freq is called when entering the op_freq production.
	EnterOp_freq(c *Op_freqContext)

	// EnterQuota_val is called when entering the quota_val production.
	EnterQuota_val(c *Quota_valContext)

	// EnterQuota_ts_node is called when entering the quota_ts_node production.
	EnterQuota_ts_node(c *Quota_ts_nodeContext)

	// EnterQuota_ts_lst is called when entering the quota_ts_lst production.
	EnterQuota_ts_lst(c *Quota_ts_lstContext)

	// EnterQuota_ts is called when entering the quota_ts production.
	EnterQuota_ts(c *Quota_tsContext)

	// EnterCreate_role_stmt is called when entering the create_role_stmt production.
	EnterCreate_role_stmt(c *Create_role_stmtContext)

	// EnterCreate_dblink_stmt is called when entering the create_dblink_stmt production.
	EnterCreate_dblink_stmt(c *Create_dblink_stmtContext)

	// EnterDb_type_str is called when entering the db_type_str production.
	EnterDb_type_str(c *Db_type_strContext)

	// EnterDblink_option_lst_options is called when entering the dblink_option_lst_options production.
	EnterDblink_option_lst_options(c *Dblink_option_lst_optionsContext)

	// EnterDblink_option_lst is called when entering the dblink_option_lst production.
	EnterDblink_option_lst(c *Dblink_option_lstContext)

	// EnterDblink_option is called when entering the dblink_option production.
	EnterDblink_option(c *Dblink_optionContext)

	// EnterCreate_synonym_stmt is called when entering the create_synonym_stmt production.
	EnterCreate_synonym_stmt(c *Create_synonym_stmtContext)

	// EnterFull_synonym_name is called when entering the full_synonym_name production.
	EnterFull_synonym_name(c *Full_synonym_nameContext)

	// EnterFull_obj_name is called when entering the full_obj_name production.
	EnterFull_obj_name(c *Full_obj_nameContext)

	// EnterCreate_domain_stmt is called when entering the create_domain_stmt production.
	EnterCreate_domain_stmt(c *Create_domain_stmtContext)

	// EnterDomain_default_option is called when entering the domain_default_option production.
	EnterDomain_default_option(c *Domain_default_optionContext)

	// EnterDomain_constraints_option is called when entering the domain_constraints_option production.
	EnterDomain_constraints_option(c *Domain_constraints_optionContext)

	// EnterDomain_constraints_def is called when entering the domain_constraints_def production.
	EnterDomain_constraints_def(c *Domain_constraints_defContext)

	// EnterDomain_constraints is called when entering the domain_constraints production.
	EnterDomain_constraints(c *Domain_constraintsContext)

	// EnterDomain_constraint is called when entering the domain_constraint production.
	EnterDomain_constraint(c *Domain_constraintContext)

	// EnterDomain_constraint_name_def_options is called when entering the domain_constraint_name_def_options production.
	EnterDomain_constraint_name_def_options(c *Domain_constraint_name_def_optionsContext)

	// EnterDomain_constraint_name_def is called when entering the domain_constraint_name_def production.
	EnterDomain_constraint_name_def(c *Domain_constraint_name_defContext)

	// EnterDomain_constraint_name is called when entering the domain_constraint_name production.
	EnterDomain_constraint_name(c *Domain_constraint_nameContext)

	// EnterCreate_character_set_stmt is called when entering the create_character_set_stmt production.
	EnterCreate_character_set_stmt(c *Create_character_set_stmtContext)

	// EnterCharacter_set_source is called when entering the character_set_source production.
	EnterCharacter_set_source(c *Character_set_sourceContext)

	// EnterExisting_character_set_name is called when entering the existing_character_set_name production.
	EnterExisting_character_set_name(c *Existing_character_set_nameContext)

	// EnterCharacter_set_name is called when entering the character_set_name production.
	EnterCharacter_set_name(c *Character_set_nameContext)

	// EnterCollate_clause_option is called when entering the collate_clause_option production.
	EnterCollate_clause_option(c *Collate_clause_optionContext)

	// EnterCollation_name is called when entering the collation_name production.
	EnterCollation_name(c *Collation_nameContext)

	// EnterCreate_collation_stmt is called when entering the create_collation_stmt production.
	EnterCreate_collation_stmt(c *Create_collation_stmtContext)

	// EnterExisting_collation_name is called when entering the existing_collation_name production.
	EnterExisting_collation_name(c *Existing_collation_nameContext)

	// EnterPad_option is called when entering the pad_option production.
	EnterPad_option(c *Pad_optionContext)

	// EnterCreate_sequence_stmt is called when entering the create_sequence_stmt production.
	EnterCreate_sequence_stmt(c *Create_sequence_stmtContext)

	// EnterSequence_option_list_options is called when entering the sequence_option_list_options production.
	EnterSequence_option_list_options(c *Sequence_option_list_optionsContext)

	// EnterSequence_option_list is called when entering the sequence_option_list production.
	EnterSequence_option_list(c *Sequence_option_listContext)

	// EnterSequence_option is called when entering the sequence_option production.
	EnterSequence_option(c *Sequence_optionContext)

	// EnterSequence_name is called when entering the sequence_name production.
	EnterSequence_name(c *Sequence_nameContext)

	// EnterIncrement_option is called when entering the increment_option production.
	EnterIncrement_option(c *Increment_optionContext)

	// EnterStart_option is called when entering the start_option production.
	EnterStart_option(c *Start_optionContext)

	// EnterCurrent_option is called when entering the current_option production.
	EnterCurrent_option(c *Current_optionContext)

	// EnterMaxvalue_option is called when entering the maxvalue_option production.
	EnterMaxvalue_option(c *Maxvalue_optionContext)

	// EnterMinvalue_option is called when entering the minvalue_option production.
	EnterMinvalue_option(c *Minvalue_optionContext)

	// EnterCycle_option is called when entering the cycle_option production.
	EnterCycle_option(c *Cycle_optionContext)

	// EnterCache_option is called when entering the cache_option production.
	EnterCache_option(c *Cache_optionContext)

	// EnterOrder_option is called when entering the order_option production.
	EnterOrder_option(c *Order_optionContext)

	// EnterSeq_local_option is called when entering the seq_local_option production.
	EnterSeq_local_option(c *Seq_local_optionContext)

	// EnterWhenever_stmt_options is called when entering the whenever_stmt_options production.
	EnterWhenever_stmt_options(c *Whenever_stmt_optionsContext)

	// EnterWhenever_stmt is called when entering the whenever_stmt production.
	EnterWhenever_stmt(c *Whenever_stmtContext)

	// EnterGrant_stmt is called when entering the grant_stmt production.
	EnterGrant_stmt(c *Grant_stmtContext)

	// EnterGrant_tag is called when entering the grant_tag production.
	EnterGrant_tag(c *Grant_tagContext)

	// EnterGrant_stmt_body is called when entering the grant_stmt_body production.
	EnterGrant_stmt_body(c *Grant_stmt_bodyContext)

	// EnterRevoke_stmt is called when entering the revoke_stmt production.
	EnterRevoke_stmt(c *Revoke_stmtContext)

	// EnterRevoke_stmt_body is called when entering the revoke_stmt_body production.
	EnterRevoke_stmt_body(c *Revoke_stmt_bodyContext)

	// EnterGrant_privs is called when entering the grant_privs production.
	EnterGrant_privs(c *Grant_privsContext)

	// EnterGrant_priv_list is called when entering the grant_priv_list production.
	EnterGrant_priv_list(c *Grant_priv_listContext)

	// EnterGrant_priv_off is called when entering the grant_priv_off production.
	EnterGrant_priv_off(c *Grant_priv_offContext)

	// EnterGrant_priv is called when entering the grant_priv production.
	EnterGrant_priv(c *Grant_privContext)

	// EnterRevoke_action is called when entering the revoke_action production.
	EnterRevoke_action(c *Revoke_actionContext)

	// EnterDb_priv_list is called when entering the db_priv_list production.
	EnterDb_priv_list(c *Db_priv_listContext)

	// EnterDb_priv is called when entering the db_priv production.
	EnterDb_priv(c *Db_privContext)

	// EnterBy_grantor is called when entering the by_grantor production.
	EnterBy_grantor(c *By_grantorContext)

	// EnterGrantees is called when entering the grantees production.
	EnterGrantees(c *GranteesContext)

	// EnterCreate_schema_stmt is called when entering the create_schema_stmt production.
	EnterCreate_schema_stmt(c *Create_schema_stmtContext)

	// EnterOprt_arg is called when entering the oprt_arg production.
	EnterOprt_arg(c *Oprt_argContext)

	// EnterOprt_arg_lst is called when entering the oprt_arg_lst production.
	EnterOprt_arg_lst(c *Oprt_arg_lstContext)

	// EnterCreate_operator_stmt is called when entering the create_operator_stmt production.
	EnterCreate_operator_stmt(c *Create_operator_stmtContext)

	// EnterDrop_operator_stmt is called when entering the drop_operator_stmt production.
	EnterDrop_operator_stmt(c *Drop_operator_stmtContext)

	// EnterGrant_and_ddl is called when entering the grant_and_ddl production.
	EnterGrant_and_ddl(c *Grant_and_ddlContext)

	// EnterTop_exp is called when entering the top_exp production.
	EnterTop_exp(c *Top_expContext)

	// EnterU_oprt is called when entering the u_oprt production.
	EnterU_oprt(c *U_oprtContext)

	// EnterQualified_u_oprt is called when entering the qualified_u_oprt production.
	EnterQualified_u_oprt(c *Qualified_u_oprtContext)

	// EnterExp_u_oprt is called when entering the exp_u_oprt production.
	EnterExp_u_oprt(c *Exp_u_oprtContext)

	// EnterRaw_exp is called when entering the raw_exp production.
	EnterRaw_exp(c *Raw_expContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterFrom_first_last_option is called when entering the from_first_last_option production.
	EnterFrom_first_last_option(c *From_first_last_optionContext)

	// EnterAfun_arg_lst is called when entering the afun_arg_lst production.
	EnterAfun_arg_lst(c *Afun_arg_lstContext)

	// EnterAfun_arg_lst_low is called when entering the afun_arg_lst_low production.
	EnterAfun_arg_lst_low(c *Afun_arg_lst_lowContext)

	// EnterIn_value_exp is called when entering the in_value_exp production.
	EnterIn_value_exp(c *In_value_expContext)

	// EnterAfun_partition_by is called when entering the afun_partition_by production.
	EnterAfun_partition_by(c *Afun_partition_byContext)

	// EnterAfun_windowing is called when entering the afun_windowing production.
	EnterAfun_windowing(c *Afun_windowingContext)

	// EnterAfun_windowing_type is called when entering the afun_windowing_type production.
	EnterAfun_windowing_type(c *Afun_windowing_typeContext)

	// EnterAfun_range_clause is called when entering the afun_range_clause production.
	EnterAfun_range_clause(c *Afun_range_clauseContext)

	// EnterPexp is called when entering the pexp production.
	EnterPexp(c *PexpContext)

	// EnterPexp_pfx is called when entering the pexp_pfx production.
	EnterPexp_pfx(c *Pexp_pfxContext)

	// EnterPexp_cast is called when entering the pexp_cast production.
	EnterPexp_cast(c *Pexp_castContext)

	// EnterPexp_b is called when entering the pexp_b production.
	EnterPexp_b(c *Pexp_bContext)

	// EnterPexp_a is called when entering the pexp_a production.
	EnterPexp_a(c *Pexp_aContext)

	// EnterPexp_c is called when entering the pexp_c production.
	EnterPexp_c(c *Pexp_cContext)

	// EnterPexp_c_insert is called when entering the pexp_c_insert production.
	EnterPexp_c_insert(c *Pexp_c_insertContext)

	// EnterPexp_d is called when entering the pexp_d production.
	EnterPexp_d(c *Pexp_dContext)

	// EnterPexp_e is called when entering the pexp_e production.
	EnterPexp_e(c *Pexp_eContext)

	// EnterPexp_pfx2 is called when entering the pexp_pfx2 production.
	EnterPexp_pfx2(c *Pexp_pfx2Context)

	// EnterMember2 is called when entering the member2 production.
	EnterMember2(c *Member2Context)

	// EnterPexp_c2_insert is called when entering the pexp_c2_insert production.
	EnterPexp_c2_insert(c *Pexp_c2_insertContext)

	// EnterMember_access2 is called when entering the member_access2 production.
	EnterMember_access2(c *Member_access2Context)

	// EnterInvocation_expression2 is called when entering the invocation_expression2 production.
	EnterInvocation_expression2(c *Invocation_expression2Context)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterMember_access is called when entering the member_access production.
	EnterMember_access(c *Member_accessContext)

	// EnterInvocation_expression is called when entering the invocation_expression production.
	EnterInvocation_expression(c *Invocation_expressionContext)

	// EnterInvocation_expression_low is called when entering the invocation_expression_low production.
	EnterInvocation_expression_low(c *Invocation_expression_lowContext)

	// EnterXmlagg_inv_expression is called when entering the xmlagg_inv_expression production.
	EnterXmlagg_inv_expression(c *Xmlagg_inv_expressionContext)

	// EnterXmlfun_inv_expression is called when entering the xmlfun_inv_expression production.
	EnterXmlfun_inv_expression(c *Xmlfun_inv_expressionContext)

	// EnterXmlele_name is called when entering the xmlele_name production.
	EnterXmlele_name(c *Xmlele_nameContext)

	// EnterXmlele_sub_lst is called when entering the xmlele_sub_lst production.
	EnterXmlele_sub_lst(c *Xmlele_sub_lstContext)

	// EnterXmlattr_lst is called when entering the xmlattr_lst production.
	EnterXmlattr_lst(c *Xmlattr_lstContext)

	// EnterXmlattr is called when entering the xmlattr production.
	EnterXmlattr(c *XmlattrContext)

	// EnterXmlval_lst is called when entering the xmlval_lst production.
	EnterXmlval_lst(c *Xmlval_lstContext)

	// EnterKeep_clause is called when entering the keep_clause production.
	EnterKeep_clause(c *Keep_clauseContext)

	// EnterWithin_clause is called when entering the within_clause production.
	EnterWithin_clause(c *Within_clauseContext)

	// EnterTypeof_expression is called when entering the typeof_expression production.
	EnterTypeof_expression(c *Typeof_expressionContext)

	// EnterNew_obj_expression is called when entering the new_obj_expression production.
	EnterNew_obj_expression(c *New_obj_expressionContext)

	// EnterNew_arr_expression is called when entering the new_arr_expression production.
	EnterNew_arr_expression(c *New_arr_expressionContext)

	// EnterArray_creation_expression is called when entering the array_creation_expression production.
	EnterArray_creation_expression(c *Array_creation_expressionContext)

	// EnterPlsql_datatype_ex is called when entering the plsql_datatype_ex production.
	EnterPlsql_datatype_ex(c *Plsql_datatype_exContext)

	// EnterNew_array_type is called when entering the new_array_type production.
	EnterNew_array_type(c *New_array_typeContext)

	// EnterOpt_array_initializer is called when entering the opt_array_initializer production.
	EnterOpt_array_initializer(c *Opt_array_initializerContext)

	// EnterArray_initializer is called when entering the array_initializer production.
	EnterArray_initializer(c *Array_initializerContext)

	// EnterVariable_initializer_list is called when entering the variable_initializer_list production.
	EnterVariable_initializer_list(c *Variable_initializer_listContext)

	// EnterVariable_initializer is called when entering the variable_initializer production.
	EnterVariable_initializer(c *Variable_initializerContext)

	// EnterOpt_comma is called when entering the opt_comma production.
	EnterOpt_comma(c *Opt_commaContext)

	// EnterSizeof_expression is called when entering the sizeof_expression production.
	EnterSizeof_expression(c *Sizeof_expressionContext)

	// EnterElement_access is called when entering the element_access production.
	EnterElement_access(c *Element_accessContext)

	// EnterDecode_case is called when entering the decode_case production.
	EnterDecode_case(c *Decode_caseContext)

	// EnterElse_exp is called when entering the else_exp production.
	EnterElse_exp(c *Else_expContext)

	// EnterBoolean_case is called when entering the boolean_case production.
	EnterBoolean_case(c *Boolean_caseContext)

	// EnterIf_exp is called when entering the if_exp production.
	EnterIf_exp(c *If_expContext)

	// EnterBool_when_list is called when entering the bool_when_list production.
	EnterBool_when_list(c *Bool_when_listContext)

	// EnterOps is called when entering the ops production.
	EnterOps(c *OpsContext)

	// EnterValue_list is called when entering the value_list production.
	EnterValue_list(c *Value_listContext)

	// EnterIn_value_list is called when entering the in_value_list production.
	EnterIn_value_list(c *In_value_listContext)

	// EnterValue_list_set is called when entering the value_list_set production.
	EnterValue_list_set(c *Value_list_setContext)

	// EnterComma_list is called when entering the comma_list production.
	EnterComma_list(c *Comma_listContext)

	// EnterIns_value_list is called when entering the ins_value_list production.
	EnterIns_value_list(c *Ins_value_listContext)

	// EnterNull_value is called when entering the null_value production.
	EnterNull_value(c *Null_valueContext)

	// EnterId_and_rsvd_word_others is called when entering the id_and_rsvd_word_others production.
	EnterId_and_rsvd_word_others(c *Id_and_rsvd_word_othersContext)

	// EnterId_and_rsvd_word is called when entering the id_and_rsvd_word production.
	EnterId_and_rsvd_word(c *Id_and_rsvd_wordContext)

	// EnterStm_param is called when entering the stm_param production.
	EnterStm_param(c *Stm_paramContext)

	// EnterStm_param_normal is called when entering the stm_param_normal production.
	EnterStm_param_normal(c *Stm_param_normalContext)

	// EnterStm_param_name is called when entering the stm_param_name production.
	EnterStm_param_name(c *Stm_param_nameContext)

	// EnterParam_name_options is called when entering the param_name_options production.
	EnterParam_name_options(c *Param_name_optionsContext)

	// EnterContains_query_exp is called when entering the contains_query_exp production.
	EnterContains_query_exp(c *Contains_query_expContext)

	// EnterContains_query_exp_lst is called when entering the contains_query_exp_lst production.
	EnterContains_query_exp_lst(c *Contains_query_exp_lstContext)

	// EnterContains_exp is called when entering the contains_exp production.
	EnterContains_exp(c *Contains_expContext)

	// EnterStrict_option is called when entering the strict_option production.
	EnterStrict_option(c *Strict_optionContext)

	// EnterWith_unique_option is called when entering the with_unique_option production.
	EnterWith_unique_option(c *With_unique_optionContext)

	// EnterType_option is called when entering the type_option production.
	EnterType_option(c *Type_optionContext)

	// EnterType_element is called when entering the type_element production.
	EnterType_element(c *Type_elementContext)

	// EnterType_element_list is called when entering the type_element_list production.
	EnterType_element_list(c *Type_element_listContext)

	// EnterBool_exp is called when entering the bool_exp production.
	EnterBool_exp(c *Bool_expContext)

	// EnterBool_exp_element is called when entering the bool_exp_element production.
	EnterBool_exp_element(c *Bool_exp_elementContext)

	// EnterQuery_any_options is called when entering the query_any_options production.
	EnterQuery_any_options(c *Query_any_optionsContext)

	// EnterGlobal_var is called when entering the global_var production.
	EnterGlobal_var(c *Global_varContext)

	// EnterReserved_word is called when entering the reserved_word production.
	EnterReserved_word(c *Reserved_wordContext)

	// EnterNew_none_reserved_word is called when entering the new_none_reserved_word production.
	EnterNew_none_reserved_word(c *New_none_reserved_wordContext)

	// EnterInterval_nresvd_word is called when entering the interval_nresvd_word production.
	EnterInterval_nresvd_word(c *Interval_nresvd_wordContext)

	// EnterVariable_resvd_word is called when entering the variable_resvd_word production.
	EnterVariable_resvd_word(c *Variable_resvd_wordContext)

	// EnterAlias_resvd_word is called when entering the alias_resvd_word production.
	EnterAlias_resvd_word(c *Alias_resvd_wordContext)

	// EnterSchname_resvd_word is called when entering the schname_resvd_word production.
	EnterSchname_resvd_word(c *Schname_resvd_wordContext)

	// EnterRaw_id is called when entering the raw_id production.
	EnterRaw_id(c *Raw_idContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterQualified_name is called when entering the qualified_name production.
	EnterQualified_name(c *Qualified_nameContext)

	// EnterQualified_name2 is called when entering the qualified_name2 production.
	EnterQualified_name2(c *Qualified_name2Context)

	// EnterVariable_name is called when entering the variable_name production.
	EnterVariable_name(c *Variable_nameContext)

	// EnterEnd_loop_label_null is called when entering the end_loop_label_null production.
	EnterEnd_loop_label_null(c *End_loop_label_nullContext)

	// EnterLabel_name_options is called when entering the label_name_options production.
	EnterLabel_name_options(c *Label_name_optionsContext)

	// EnterLabel_name is called when entering the label_name production.
	EnterLabel_name(c *Label_nameContext)

	// EnterDatabase_name is called when entering the database_name production.
	EnterDatabase_name(c *Database_nameContext)

	// EnterBackup_name is called when entering the backup_name production.
	EnterBackup_name(c *Backup_nameContext)

	// EnterFull_proc_name is called when entering the full_proc_name production.
	EnterFull_proc_name(c *Full_proc_nameContext)

	// EnterFull_proc_name2 is called when entering the full_proc_name2 production.
	EnterFull_proc_name2(c *Full_proc_name2Context)

	// EnterFull_fun_name is called when entering the full_fun_name production.
	EnterFull_fun_name(c *Full_fun_nameContext)

	// EnterFull_table_name is called when entering the full_table_name production.
	EnterFull_table_name(c *Full_table_nameContext)

	// EnterFull_grp_name is called when entering the full_grp_name production.
	EnterFull_grp_name(c *Full_grp_nameContext)

	// EnterFull_table_name2 is called when entering the full_table_name2 production.
	EnterFull_table_name2(c *Full_table_name2Context)

	// EnterFull_partition_name is called when entering the full_partition_name production.
	EnterFull_partition_name(c *Full_partition_nameContext)

	// EnterFull_schema_name is called when entering the full_schema_name production.
	EnterFull_schema_name(c *Full_schema_nameContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterConstraint_name is called when entering the constraint_name production.
	EnterConstraint_name(c *Constraint_nameContext)

	// EnterFull_trigger_name is called when entering the full_trigger_name production.
	EnterFull_trigger_name(c *Full_trigger_nameContext)

	// EnterFull_trigger_name2 is called when entering the full_trigger_name2 production.
	EnterFull_trigger_name2(c *Full_trigger_name2Context)

	// EnterFull_view_name is called when entering the full_view_name production.
	EnterFull_view_name(c *Full_view_nameContext)

	// EnterFull_view_name2 is called when entering the full_view_name2 production.
	EnterFull_view_name2(c *Full_view_name2Context)

	// EnterCursor_name is called when entering the cursor_name production.
	EnterCursor_name(c *Cursor_nameContext)

	// EnterTrigger_name is called when entering the trigger_name production.
	EnterTrigger_name(c *Trigger_nameContext)

	// EnterLogin_name is called when entering the login_name production.
	EnterLogin_name(c *Login_nameContext)

	// EnterProfile_name is called when entering the profile_name production.
	EnterProfile_name(c *Profile_nameContext)

	// EnterUser_name is called when entering the user_name production.
	EnterUser_name(c *User_nameContext)

	// EnterRole_name is called when entering the role_name production.
	EnterRole_name(c *Role_nameContext)

	// EnterUser_role_name is called when entering the user_role_name production.
	EnterUser_role_name(c *User_role_nameContext)

	// EnterRole_name_list is called when entering the role_name_list production.
	EnterRole_name_list(c *Role_name_listContext)

	// EnterFull_func_name is called when entering the full_func_name production.
	EnterFull_func_name(c *Full_func_nameContext)

	// EnterParam_name is called when entering the param_name production.
	EnterParam_name(c *Param_nameContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterIndex_name2 is called when entering the index_name2 production.
	EnterIndex_name2(c *Index_name2Context)

	// EnterTrig_old_name is called when entering the trig_old_name production.
	EnterTrig_old_name(c *Trig_old_nameContext)

	// EnterTrig_new_name is called when entering the trig_new_name production.
	EnterTrig_new_name(c *Trig_new_nameContext)

	// EnterFull_tv_name is called when entering the full_tv_name production.
	EnterFull_tv_name(c *Full_tv_nameContext)

	// EnterFull_object_name is called when entering the full_object_name production.
	EnterFull_object_name(c *Full_object_nameContext)

	// EnterOrient_option is called when entering the orient_option production.
	EnterOrient_option(c *Orient_optionContext)

	// EnterDatepart is called when entering the datepart production.
	EnterDatepart(c *DatepartContext)

	// EnterDatepart_op is called when entering the datepart_op production.
	EnterDatepart_op(c *Datepart_opContext)

	// EnterDatead_fun is called when entering the datead_fun production.
	EnterDatead_fun(c *Datead_funContext)

	// EnterReturning is called when entering the returning production.
	EnterReturning(c *ReturningContext)

	// EnterPretty is called when entering the pretty production.
	EnterPretty(c *PrettyContext)

	// EnterWrapper_flag is called when entering the wrapper_flag production.
	EnterWrapper_flag(c *Wrapper_flagContext)

	// EnterArray_wrapper is called when entering the array_wrapper production.
	EnterArray_wrapper(c *Array_wrapperContext)

	// EnterJson_tail_on_empty is called when entering the json_tail_on_empty production.
	EnterJson_tail_on_empty(c *Json_tail_on_emptyContext)

	// EnterEmpty_handle is called when entering the empty_handle production.
	EnterEmpty_handle(c *Empty_handleContext)

	// EnterJson_tail_on_error_null is called when entering the json_tail_on_error_null production.
	EnterJson_tail_on_error_null(c *Json_tail_on_error_nullContext)

	// EnterError_handle is called when entering the error_handle production.
	EnterError_handle(c *Error_handleContext)

	// EnterSavepoint_name is called when entering the savepoint_name production.
	EnterSavepoint_name(c *Savepoint_nameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterAlias_2 is called when entering the alias_2 production.
	EnterAlias_2(c *Alias_2Context)

	// EnterFull_column_name is called when entering the full_column_name production.
	EnterFull_column_name(c *Full_column_nameContext)

	// EnterSchema_name is called when entering the schema_name production.
	EnterSchema_name(c *Schema_nameContext)

	// EnterNot_tag is called when entering the not_tag production.
	EnterNot_tag(c *Not_tagContext)

	// EnterDebug_tag is called when entering the debug_tag production.
	EnterDebug_tag(c *Debug_tagContext)

	// EnterColumn_tag is called when entering the column_tag production.
	EnterColumn_tag(c *Column_tagContext)

	// EnterPendant_tag is called when entering the pendant_tag production.
	EnterPendant_tag(c *Pendant_tagContext)

	// EnterUnique_tag is called when entering the unique_tag production.
	EnterUnique_tag(c *Unique_tagContext)

	// EnterPartition_tag is called when entering the partition_tag production.
	EnterPartition_tag(c *Partition_tagContext)

	// EnterRow_tag is called when entering the row_tag production.
	EnterRow_tag(c *Row_tagContext)

	// EnterAs_tag is called when entering the as_tag production.
	EnterAs_tag(c *As_tagContext)

	// EnterFrom_tag is called when entering the from_tag production.
	EnterFrom_tag(c *From_tagContext)

	// EnterInto_tag is called when entering the into_tag production.
	EnterInto_tag(c *Into_tagContext)

	// EnterWork_tag is called when entering the work_tag production.
	EnterWork_tag(c *Work_tagContext)

	// EnterWith_grant_option is called when entering the with_grant_option production.
	EnterWith_grant_option(c *With_grant_optionContext)

	// EnterWith_admin_option is called when entering the with_admin_option production.
	EnterWith_admin_option(c *With_admin_optionContext)

	// EnterTime_zone_or_local is called when entering the time_zone_or_local production.
	EnterTime_zone_or_local(c *Time_zone_or_localContext)

	// EnterSub_plsql_datatype is called when entering the sub_plsql_datatype production.
	EnterSub_plsql_datatype(c *Sub_plsql_datatypeContext)

	// EnterDatatype_list is called when entering the datatype_list production.
	EnterDatatype_list(c *Datatype_listContext)

	// EnterDatatype is called when entering the datatype production.
	EnterDatatype(c *DatatypeContext)

	// EnterDatatype2 is called when entering the datatype2 production.
	EnterDatatype2(c *Datatype2Context)

	// EnterOpr_dtype is called when entering the opr_dtype production.
	EnterOpr_dtype(c *Opr_dtypeContext)

	// EnterOpr_datatype_lst is called when entering the opr_datatype_lst production.
	EnterOpr_datatype_lst(c *Opr_datatype_lstContext)

	// EnterInterval_qualifier is called when entering the interval_qualifier production.
	EnterInterval_qualifier(c *Interval_qualifierContext)

	// EnterDtype is called when entering the dtype production.
	EnterDtype(c *DtypeContext)

	// EnterDtype1 is called when entering the dtype1 production.
	EnterDtype1(c *Dtype1Context)

	// EnterDtype2 is called when entering the dtype2 production.
	EnterDtype2(c *Dtype2Context)

	// EnterDouble_length_option is called when entering the double_length_option production.
	EnterDouble_length_option(c *Double_length_optionContext)

	// EnterSize_unit_caluse is called when entering the size_unit_caluse production.
	EnterSize_unit_caluse(c *Size_unit_caluseContext)

	// EnterLt_integer_negative is called when entering the lt_integer_negative production.
	EnterLt_integer_negative(c *Lt_integer_negativeContext)

	// EnterCreate_contextindex_stmt is called when entering the create_contextindex_stmt production.
	EnterCreate_contextindex_stmt(c *Create_contextindex_stmtContext)

	// EnterLexer_name is called when entering the lexer_name production.
	EnterLexer_name(c *Lexer_nameContext)

	// EnterLexer_clause is called when entering the lexer_clause production.
	EnterLexer_clause(c *Lexer_clauseContext)

	// EnterLexer_clause2 is called when entering the lexer_clause2 production.
	EnterLexer_clause2(c *Lexer_clause2Context)

	// EnterSync is called when entering the sync production.
	EnterSync(c *SyncContext)

	// EnterDrop_contextindex_stmt is called when entering the drop_contextindex_stmt production.
	EnterDrop_contextindex_stmt(c *Drop_contextindex_stmtContext)

	// EnterAlter_contextindex_stmt is called when entering the alter_contextindex_stmt production.
	EnterAlter_contextindex_stmt(c *Alter_contextindex_stmtContext)

	// EnterCti_sync_option is called when entering the cti_sync_option production.
	EnterCti_sync_option(c *Cti_sync_optionContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterSizeof_type is called when entering the sizeof_type production.
	EnterSizeof_type(c *Sizeof_typeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterArray_type is called when entering the array_type production.
	EnterArray_type(c *Array_typeContext)

	// EnterBuiltin_types is called when entering the builtin_types production.
	EnterBuiltin_types(c *Builtin_typesContext)

	// EnterIntegral_type is called when entering the integral_type production.
	EnterIntegral_type(c *Integral_typeContext)

	// EnterSql_builtin_types is called when entering the sql_builtin_types production.
	EnterSql_builtin_types(c *Sql_builtin_typesContext)

	// EnterCursor_declaration is called when entering the cursor_declaration production.
	EnterCursor_declaration(c *Cursor_declarationContext)

	// EnterCursor_declaration_2 is called when entering the cursor_declaration_2 production.
	EnterCursor_declaration_2(c *Cursor_declaration_2Context)

	// EnterCursor_attrs_options is called when entering the cursor_attrs_options production.
	EnterCursor_attrs_options(c *Cursor_attrs_optionsContext)

	// EnterCursor_attrs is called when entering the cursor_attrs production.
	EnterCursor_attrs(c *Cursor_attrsContext)

	// EnterCursor_attr is called when entering the cursor_attr production.
	EnterCursor_attr(c *Cursor_attrContext)

	// EnterOpt_rank_specifier is called when entering the opt_rank_specifier production.
	EnterOpt_rank_specifier(c *Opt_rank_specifierContext)

	// EnterRank_specifiers is called when entering the rank_specifiers production.
	EnterRank_specifiers(c *Rank_specifiersContext)

	// EnterRank_specifier is called when entering the rank_specifier production.
	EnterRank_specifier(c *Rank_specifierContext)

	// EnterOpt_dim_separators is called when entering the opt_dim_separators production.
	EnterOpt_dim_separators(c *Opt_dim_separatorsContext)

	// EnterOpt_rank_specifier2 is called when entering the opt_rank_specifier2 production.
	EnterOpt_rank_specifier2(c *Opt_rank_specifier2Context)

	// EnterDim_separators is called when entering the dim_separators production.
	EnterDim_separators(c *Dim_separatorsContext)

	// EnterOpt_argument_list is called when entering the opt_argument_list production.
	EnterOpt_argument_list(c *Opt_argument_listContext)

	// EnterJson_fun_tail is called when entering the json_fun_tail production.
	EnterJson_fun_tail(c *Json_fun_tailContext)

	// EnterIgnore_nulls_clause is called when entering the ignore_nulls_clause production.
	EnterIgnore_nulls_clause(c *Ignore_nulls_clauseContext)

	// EnterMixed_param_list is called when entering the mixed_param_list production.
	EnterMixed_param_list(c *Mixed_param_listContext)

	// EnterMixed_param is called when entering the mixed_param production.
	EnterMixed_param(c *Mixed_paramContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterCursor_option is called when entering the cursor_option production.
	EnterCursor_option(c *Cursor_optionContext)

	// EnterWithout_into_select2 is called when entering the without_into_select2 production.
	EnterWithout_into_select2(c *Without_into_select2Context)

	// EnterCursor_option_2 is called when entering the cursor_option_2 production.
	EnterCursor_option_2(c *Cursor_option_2Context)

	// EnterRegion_size is called when entering the region_size production.
	EnterRegion_size(c *Region_sizeContext)

	// EnterCopy_num is called when entering the copy_num production.
	EnterCopy_num(c *Copy_numContext)

	// EnterRedundancy_clause is called when entering the redundancy_clause production.
	EnterRedundancy_clause(c *Redundancy_clauseContext)

	// EnterStriping_clause is called when entering the striping_clause production.
	EnterStriping_clause(c *Striping_clauseContext)

	// EnterWith_huge_clause is called when entering the with_huge_clause production.
	EnterWith_huge_clause(c *With_huge_clauseContext)

	// ExitDmprogram is called when exiting the dmprogram production.
	ExitDmprogram(c *DmprogramContext)

	// ExitSql_clauses is called when exiting the sql_clauses production.
	ExitSql_clauses(c *Sql_clausesContext)

	// ExitDdlsql is called when exiting the ddlsql production.
	ExitDdlsql(c *DdlsqlContext)

	// ExitDmlsql is called when exiting the dmlsql production.
	ExitDmlsql(c *DmlsqlContext)

	// ExitPrivsql is called when exiting the privsql production.
	ExitPrivsql(c *PrivsqlContext)

	// ExitOthersql is called when exiting the othersql production.
	ExitOthersql(c *OthersqlContext)

	// ExitUtilsql is called when exiting the utilsql production.
	ExitUtilsql(c *UtilsqlContext)

	// ExitExplainsql is called when exiting the explainsql production.
	ExitExplainsql(c *ExplainsqlContext)

	// ExitShutdown_stmt is called when exiting the shutdown_stmt production.
	ExitShutdown_stmt(c *Shutdown_stmtContext)

	// ExitAlter_diskgroup_stmt is called when exiting the alter_diskgroup_stmt production.
	ExitAlter_diskgroup_stmt(c *Alter_diskgroup_stmtContext)

	// ExitLocal is called when exiting the local production.
	ExitLocal(c *LocalContext)

	// ExitDmsubprogram is called when exiting the dmsubprogram production.
	ExitDmsubprogram(c *DmsubprogramContext)

	// ExitDeclare_block is called when exiting the declare_block production.
	ExitDeclare_block(c *Declare_blockContext)

	// ExitDecl_var_cur_list_options is called when exiting the decl_var_cur_list_options production.
	ExitDecl_var_cur_list_options(c *Decl_var_cur_list_optionsContext)

	// ExitDecl_var_cur_list_1 is called when exiting the decl_var_cur_list_1 production.
	ExitDecl_var_cur_list_1(c *Decl_var_cur_list_1Context)

	// ExitDecl_var_cur_list is called when exiting the decl_var_cur_list production.
	ExitDecl_var_cur_list(c *Decl_var_cur_listContext)

	// ExitDecl_plsql_type is called when exiting the decl_plsql_type production.
	ExitDecl_plsql_type(c *Decl_plsql_typeContext)

	// ExitPlsql_type_def is called when exiting the plsql_type_def production.
	ExitPlsql_type_def(c *Plsql_type_defContext)

	// ExitLt_int_lst is called when exiting the lt_int_lst production.
	ExitLt_int_lst(c *Lt_int_lstContext)

	// ExitRec_item_def_list is called when exiting the rec_item_def_list production.
	ExitRec_item_def_list(c *Rec_item_def_listContext)

	// ExitRec_item_def is called when exiting the rec_item_def production.
	ExitRec_item_def(c *Rec_item_defContext)

	// ExitDecl_variable is called when exiting the decl_variable production.
	ExitDecl_variable(c *Decl_variableContext)

	// ExitNot_null is called when exiting the not_null production.
	ExitNot_null(c *Not_nullContext)

	// ExitPlsql_datatype is called when exiting the plsql_datatype production.
	ExitPlsql_datatype(c *Plsql_datatypeContext)

	// ExitDefault_clause_option is called when exiting the default_clause_option production.
	ExitDefault_clause_option(c *Default_clause_optionContext)

	// ExitVariable_name_list is called when exiting the variable_name_list production.
	ExitVariable_name_list(c *Variable_name_listContext)

	// ExitDecl_except is called when exiting the decl_except production.
	ExitDecl_except(c *Decl_exceptContext)

	// ExitPragma_def is called when exiting the pragma_def production.
	ExitPragma_def(c *Pragma_defContext)

	// ExitPragma is called when exiting the pragma production.
	ExitPragma(c *PragmaContext)

	// ExitPlbody is called when exiting the plbody production.
	ExitPlbody(c *PlbodyContext)

	// ExitSs_plbody is called when exiting the ss_plbody production.
	ExitSs_plbody(c *Ss_plbodyContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitLabel_list is called when exiting the label_list production.
	ExitLabel_list(c *Label_listContext)

	// ExitLabel_list_options is called when exiting the label_list_options production.
	ExitLabel_list_options(c *Label_list_optionsContext)

	// ExitLabel_demiliter_l is called when exiting the label_demiliter_l production.
	ExitLabel_demiliter_l(c *Label_demiliter_lContext)

	// ExitLabel_demiliter_r is called when exiting the label_demiliter_r production.
	ExitLabel_demiliter_r(c *Label_demiliter_rContext)

	// ExitPlsql is called when exiting the plsql production.
	ExitPlsql(c *PlsqlContext)

	// ExitUr_option is called when exiting the ur_option production.
	ExitUr_option(c *Ur_optionContext)

	// ExitFlashback_trig_enable is called when exiting the flashback_trig_enable production.
	ExitFlashback_trig_enable(c *Flashback_trig_enableContext)

	// ExitScn_or_lsn is called when exiting the scn_or_lsn production.
	ExitScn_or_lsn(c *Scn_or_lsnContext)

	// ExitFull_table_name_list is called when exiting the full_table_name_list production.
	ExitFull_table_name_list(c *Full_table_name_listContext)

	// ExitFlashback_tab_stmt is called when exiting the flashback_tab_stmt production.
	ExitFlashback_tab_stmt(c *Flashback_tab_stmtContext)

	// ExitRename is called when exiting the rename production.
	ExitRename(c *RenameContext)

	// ExitAlter_system_set_stmt is called when exiting the alter_system_set_stmt production.
	ExitAlter_system_set_stmt(c *Alter_system_set_stmtContext)

	// ExitDefer is called when exiting the defer production.
	ExitDefer(c *DeferContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

	// ExitAlter_session_stmt is called when exiting the alter_session_stmt production.
	ExitAlter_session_stmt(c *Alter_session_stmtContext)

	// ExitParallel_mode is called when exiting the parallel_mode production.
	ExitParallel_mode(c *Parallel_modeContext)

	// ExitParallel_degree is called when exiting the parallel_degree production.
	ExitParallel_degree(c *Parallel_degreeContext)

	// ExitPurge is called when exiting the purge production.
	ExitPurge(c *PurgeContext)

	// ExitSess_id is called when exiting the sess_id production.
	ExitSess_id(c *Sess_idContext)

	// ExitSet_time_zone_string is called when exiting the set_time_zone_string production.
	ExitSet_time_zone_string(c *Set_time_zone_stringContext)

	// ExitSess_attr is called when exiting the sess_attr production.
	ExitSess_attr(c *Sess_attrContext)

	// ExitSess_attr_val is called when exiting the sess_attr_val production.
	ExitSess_attr_val(c *Sess_attr_valContext)

	// ExitSet_schema_stmt is called when exiting the set_schema_stmt production.
	ExitSet_schema_stmt(c *Set_schema_stmtContext)

	// ExitPlblock is called when exiting the plblock production.
	ExitPlblock(c *PlblockContext)

	// ExitExcept_option is called when exiting the except_option production.
	ExitExcept_option(c *Except_optionContext)

	// ExitFinally_option is called when exiting the finally_option production.
	ExitFinally_option(c *Finally_optionContext)

	// ExitFinally_tail is called when exiting the finally_tail production.
	ExitFinally_tail(c *Finally_tailContext)

	// ExitExcept_handler_list is called when exiting the except_handler_list production.
	ExitExcept_handler_list(c *Except_handler_listContext)

	// ExitExcept_handler is called when exiting the except_handler production.
	ExitExcept_handler(c *Except_handlerContext)

	// ExitExcept_name is called when exiting the except_name production.
	ExitExcept_name(c *Except_nameContext)

	// ExitExcept_list is called when exiting the except_list production.
	ExitExcept_list(c *Except_listContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitIf_stmt_clause is called when exiting the if_stmt_clause production.
	ExitIf_stmt_clause(c *If_stmt_clauseContext)

	// ExitIf_condition_clause is called when exiting the if_condition_clause production.
	ExitIf_condition_clause(c *If_condition_clauseContext)

	// ExitIf_then_clause is called when exiting the if_then_clause production.
	ExitIf_then_clause(c *If_then_clauseContext)

	// ExitElseif_lst_option is called when exiting the elseif_lst_option production.
	ExitElseif_lst_option(c *Elseif_lst_optionContext)

	// ExitElseif_clause is called when exiting the elseif_clause production.
	ExitElseif_clause(c *Elseif_clauseContext)

	// ExitElse_option is called when exiting the else_option production.
	ExitElse_option(c *Else_optionContext)

	// ExitSs_if_stmt_clause is called when exiting the ss_if_stmt_clause production.
	ExitSs_if_stmt_clause(c *Ss_if_stmt_clauseContext)

	// ExitSs_elseif_lst_option is called when exiting the ss_elseif_lst_option production.
	ExitSs_elseif_lst_option(c *Ss_elseif_lst_optionContext)

	// ExitSs_elseif_clause is called when exiting the ss_elseif_clause production.
	ExitSs_elseif_clause(c *Ss_elseif_clauseContext)

	// ExitSs_else_option is called when exiting the ss_else_option production.
	ExitSs_else_option(c *Ss_else_optionContext)

	// ExitCase_stmt is called when exiting the case_stmt production.
	ExitCase_stmt(c *Case_stmtContext)

	// ExitPlsearched_when_clause is called when exiting the plsearched_when_clause production.
	ExitPlsearched_when_clause(c *Plsearched_when_clauseContext)

	// ExitPlsearched_when_list is called when exiting the plsearched_when_list production.
	ExitPlsearched_when_list(c *Plsearched_when_listContext)

	// ExitCase_option is called when exiting the case_option production.
	ExitCase_option(c *Case_optionContext)

	// ExitAssign_stmt is called when exiting the assign_stmt production.
	ExitAssign_stmt(c *Assign_stmtContext)

	// ExitAssign_obj is called when exiting the assign_obj production.
	ExitAssign_obj(c *Assign_objContext)

	// ExitAssign_obj2 is called when exiting the assign_obj2 production.
	ExitAssign_obj2(c *Assign_obj2Context)

	// ExitAssign_op is called when exiting the assign_op production.
	ExitAssign_op(c *Assign_opContext)

	// ExitGoto_stmt is called when exiting the goto_stmt production.
	ExitGoto_stmt(c *Goto_stmtContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitLoop_stmt is called when exiting the loop_stmt production.
	ExitLoop_stmt(c *Loop_stmtContext)

	// ExitRepeat_stmt is called when exiting the repeat_stmt production.
	ExitRepeat_stmt(c *Repeat_stmtContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitForall_stmt is called when exiting the forall_stmt production.
	ExitForall_stmt(c *Forall_stmtContext)

	// ExitForall_between_option is called when exiting the forall_between_option production.
	ExitForall_between_option(c *Forall_between_optionContext)

	// ExitForall_save_exception_option is called when exiting the forall_save_exception_option production.
	ExitForall_save_exception_option(c *Forall_save_exception_optionContext)

	// ExitForall_index_values is called when exiting the forall_index_values production.
	ExitForall_index_values(c *Forall_index_valuesContext)

	// ExitForall_start is called when exiting the forall_start production.
	ExitForall_start(c *Forall_startContext)

	// ExitForall_dml_stmt is called when exiting the forall_dml_stmt production.
	ExitForall_dml_stmt(c *Forall_dml_stmtContext)

	// ExitIn_option is called when exiting the in_option production.
	ExitIn_option(c *In_optionContext)

	// ExitFor_condition is called when exiting the for_condition production.
	ExitFor_condition(c *For_conditionContext)

	// ExitPipe_row_stmt is called when exiting the pipe_row_stmt production.
	ExitPipe_row_stmt(c *Pipe_row_stmtContext)

	// ExitExit_stmt is called when exiting the exit_stmt production.
	ExitExit_stmt(c *Exit_stmtContext)

	// ExitContinue_stmt is called when exiting the continue_stmt production.
	ExitContinue_stmt(c *Continue_stmtContext)

	// ExitNull_stmt is called when exiting the null_stmt production.
	ExitNull_stmt(c *Null_stmtContext)

	// ExitPrint_stmt is called when exiting the print_stmt production.
	ExitPrint_stmt(c *Print_stmtContext)

	// ExitExecute_stmt is called when exiting the execute_stmt production.
	ExitExecute_stmt(c *Execute_stmtContext)

	// ExitDyn_return is called when exiting the dyn_return production.
	ExitDyn_return(c *Dyn_returnContext)

	// ExitUsing_clause is called when exiting the using_clause production.
	ExitUsing_clause(c *Using_clauseContext)

	// ExitUsing_exp_list is called when exiting the using_exp_list production.
	ExitUsing_exp_list(c *Using_exp_listContext)

	// ExitUsing_exp is called when exiting the using_exp production.
	ExitUsing_exp(c *Using_expContext)

	// ExitAlter_proc_stmt is called when exiting the alter_proc_stmt production.
	ExitAlter_proc_stmt(c *Alter_proc_stmtContext)

	// ExitAlter_func_stmt is called when exiting the alter_func_stmt production.
	ExitAlter_func_stmt(c *Alter_func_stmtContext)

	// ExitAlter_package_stmt is called when exiting the alter_package_stmt production.
	ExitAlter_package_stmt(c *Alter_package_stmtContext)

	// ExitPkg_type is called when exiting the pkg_type production.
	ExitPkg_type(c *Pkg_typeContext)

	// ExitDeclare_opt is called when exiting the declare_opt production.
	ExitDeclare_opt(c *Declare_optContext)

	// ExitAlter_table_stmt is called when exiting the alter_table_stmt production.
	ExitAlter_table_stmt(c *Alter_table_stmtContext)

	// ExitAlter_tag is called when exiting the alter_tag production.
	ExitAlter_tag(c *Alter_tagContext)

	// ExitAlter_index_stmt is called when exiting the alter_index_stmt production.
	ExitAlter_index_stmt(c *Alter_index_stmtContext)

	// ExitFull_index_name is called when exiting the full_index_name production.
	ExitFull_index_name(c *Full_index_nameContext)

	// ExitAlter_index_action is called when exiting the alter_index_action production.
	ExitAlter_index_action(c *Alter_index_actionContext)

	// ExitRebuild_clause is called when exiting the rebuild_clause production.
	ExitRebuild_clause(c *Rebuild_clauseContext)

	// ExitExclusive_options is called when exiting the exclusive_options production.
	ExitExclusive_options(c *Exclusive_optionsContext)

	// ExitAsynchronous_options is called when exiting the asynchronous_options production.
	ExitAsynchronous_options(c *Asynchronous_optionsContext)

	// ExitVisible_clause is called when exiting the visible_clause production.
	ExitVisible_clause(c *Visible_clauseContext)

	// ExitColumn_def_list is called when exiting the column_def_list production.
	ExitColumn_def_list(c *Column_def_listContext)

	// ExitLock is called when exiting the lock production.
	ExitLock(c *LockContext)

	// ExitAlter_table_partition_action is called when exiting the alter_table_partition_action production.
	ExitAlter_table_partition_action(c *Alter_table_partition_actionContext)

	// ExitTemplate_info is called when exiting the template_info production.
	ExitTemplate_info(c *Template_infoContext)

	// ExitTemplate_item_2 is called when exiting the template_item_2 production.
	ExitTemplate_item_2(c *Template_item_2Context)

	// ExitTemplate_item_1 is called when exiting the template_item_1 production.
	ExitTemplate_item_1(c *Template_item_1Context)

	// ExitIncluding_indexes is called when exiting the including_indexes production.
	ExitIncluding_indexes(c *Including_indexesContext)

	// ExitTruncate_partition_name is called when exiting the truncate_partition_name production.
	ExitTruncate_partition_name(c *Truncate_partition_nameContext)

	// ExitCons_enable is called when exiting the cons_enable production.
	ExitCons_enable(c *Cons_enableContext)

	// ExitReuse_storage_option is called when exiting the reuse_storage_option production.
	ExitReuse_storage_option(c *Reuse_storage_optionContext)

	// ExitAlter_table_action is called when exiting the alter_table_action production.
	ExitAlter_table_action(c *Alter_table_actionContext)

	// ExitFast_flag is called when exiting the fast_flag production.
	ExitFast_flag(c *Fast_flagContext)

	// ExitStorage_stat_flag is called when exiting the storage_stat_flag production.
	ExitStorage_stat_flag(c *Storage_stat_flagContext)

	// ExitStorage_stat_cols is called when exiting the storage_stat_cols production.
	ExitStorage_stat_cols(c *Storage_stat_colsContext)

	// ExitHfs_rebuild_level is called when exiting the hfs_rebuild_level production.
	ExitHfs_rebuild_level(c *Hfs_rebuild_levelContext)

	// ExitAta_lock_option is called when exiting the ata_lock_option production.
	ExitAta_lock_option(c *Ata_lock_optionContext)

	// ExitMdf_column_def_list is called when exiting the mdf_column_def_list production.
	ExitMdf_column_def_list(c *Mdf_column_def_listContext)

	// ExitMdf_column_def is called when exiting the mdf_column_def production.
	ExitMdf_column_def(c *Mdf_column_defContext)

	// ExitColumn_def is called when exiting the column_def production.
	ExitColumn_def(c *Column_defContext)

	// ExitColumn_def_ex is called when exiting the column_def_ex production.
	ExitColumn_def_ex(c *Column_def_exContext)

	// ExitColumn_def_low is called when exiting the column_def_low production.
	ExitColumn_def_low(c *Column_def_lowContext)

	// ExitVirtual_column_datatype is called when exiting the virtual_column_datatype production.
	ExitVirtual_column_datatype(c *Virtual_column_datatypeContext)

	// ExitVirtual_column_generated is called when exiting the virtual_column_generated production.
	ExitVirtual_column_generated(c *Virtual_column_generatedContext)

	// ExitVirtual_column_virtual is called when exiting the virtual_column_virtual production.
	ExitVirtual_column_virtual(c *Virtual_column_virtualContext)

	// ExitVirtual_column_visible is called when exiting the virtual_column_visible production.
	ExitVirtual_column_visible(c *Virtual_column_visibleContext)

	// ExitVirtual_column_def is called when exiting the virtual_column_def production.
	ExitVirtual_column_def(c *Virtual_column_defContext)

	// ExitCharset_option is called when exiting the charset_option production.
	ExitCharset_option(c *Charset_optionContext)

	// ExitColumn_def_4_option is called when exiting the column_def_4_option production.
	ExitColumn_def_4_option(c *Column_def_4_optionContext)

	// ExitAuto_update_clause is called when exiting the auto_update_clause production.
	ExitAuto_update_clause(c *Auto_update_clauseContext)

	// ExitUpdate_exp is called when exiting the update_exp production.
	ExitUpdate_exp(c *Update_expContext)

	// ExitIdentity_clause is called when exiting the identity_clause production.
	ExitIdentity_clause(c *Identity_clauseContext)

	// ExitDefault_clause_with_on_null_opt is called when exiting the default_clause_with_on_null_opt production.
	ExitDefault_clause_with_on_null_opt(c *Default_clause_with_on_null_optContext)

	// ExitDefault_clause is called when exiting the default_clause production.
	ExitDefault_clause(c *Default_clauseContext)

	// ExitDefault_exp is called when exiting the default_exp production.
	ExitDefault_exp(c *Default_expContext)

	// ExitColumn_constraint_def is called when exiting the column_constraint_def production.
	ExitColumn_constraint_def(c *Column_constraint_defContext)

	// ExitConstraint_name_def_options is called when exiting the constraint_name_def_options production.
	ExitConstraint_name_def_options(c *Constraint_name_def_optionsContext)

	// ExitConstraint_name_def is called when exiting the constraint_name_def production.
	ExitConstraint_name_def(c *Constraint_name_defContext)

	// ExitColumn_constraints is called when exiting the column_constraints production.
	ExitColumn_constraints(c *Column_constraintsContext)

	// ExitColumn_constraint is called when exiting the column_constraint production.
	ExitColumn_constraint(c *Column_constraintContext)

	// ExitColumn_constraint_action is called when exiting the column_constraint_action production.
	ExitColumn_constraint_action(c *Column_constraint_actionContext)

	// ExitNot_null_spec is called when exiting the not_null_spec production.
	ExitNot_null_spec(c *Not_null_specContext)

	// ExitUnique_spec is called when exiting the unique_spec production.
	ExitUnique_spec(c *Unique_specContext)

	// ExitRefs_spec is called when exiting the refs_spec production.
	ExitRefs_spec(c *Refs_specContext)

	// ExitRefs_spec_action is called when exiting the refs_spec_action production.
	ExitRefs_spec_action(c *Refs_spec_actionContext)

	// ExitForeign_key is called when exiting the foreign_key production.
	ExitForeign_key(c *Foreign_keyContext)

	// ExitRefd_table_and_columns is called when exiting the refd_table_and_columns production.
	ExitRefd_table_and_columns(c *Refd_table_and_columnsContext)

	// ExitRef_column_list is called when exiting the ref_column_list production.
	ExitRef_column_list(c *Ref_column_listContext)

	// ExitColumn_list is called when exiting the column_list production.
	ExitColumn_list(c *Column_listContext)

	// ExitColumn_list2 is called when exiting the column_list2 production.
	ExitColumn_list2(c *Column_list2Context)

	// ExitFull_column_list is called when exiting the full_column_list production.
	ExitFull_column_list(c *Full_column_listContext)

	// ExitColumn_list_list is called when exiting the column_list_list production.
	ExitColumn_list_list(c *Column_list_listContext)

	// ExitDrop_column_list is called when exiting the drop_column_list production.
	ExitDrop_column_list(c *Drop_column_listContext)

	// ExitMatch_option is called when exiting the match_option production.
	ExitMatch_option(c *Match_optionContext)

	// ExitMatch_type is called when exiting the match_type production.
	ExitMatch_type(c *Match_typeContext)

	// ExitRef_triggered_action is called when exiting the ref_triggered_action production.
	ExitRef_triggered_action(c *Ref_triggered_actionContext)

	// ExitUpdate_rule is called when exiting the update_rule production.
	ExitUpdate_rule(c *Update_ruleContext)

	// ExitDelete_rule is called when exiting the delete_rule production.
	ExitDelete_rule(c *Delete_ruleContext)

	// ExitRef_action is called when exiting the ref_action production.
	ExitRef_action(c *Ref_actionContext)

	// ExitCheck_constraint_def is called when exiting the check_constraint_def production.
	ExitCheck_constraint_def(c *Check_constraint_defContext)

	// ExitCheck_condition is called when exiting the check_condition production.
	ExitCheck_condition(c *Check_conditionContext)

	// ExitRestrict_cascade is called when exiting the restrict_cascade production.
	ExitRestrict_cascade(c *Restrict_cascadeContext)

	// ExitCascade_opt is called when exiting the cascade_opt production.
	ExitCascade_opt(c *Cascade_optContext)

	// ExitConstraint_name_options is called when exiting the constraint_name_options production.
	ExitConstraint_name_options(c *Constraint_name_optionsContext)

	// ExitCheck_option_def_true is called when exiting the check_option_def_true production.
	ExitCheck_option_def_true(c *Check_option_def_trueContext)

	// ExitConstraint_attributes_options is called when exiting the constraint_attributes_options production.
	ExitConstraint_attributes_options(c *Constraint_attributes_optionsContext)

	// ExitConstraint_attributes is called when exiting the constraint_attributes production.
	ExitConstraint_attributes(c *Constraint_attributesContext)

	// ExitDeferrable_option is called when exiting the deferrable_option production.
	ExitDeferrable_option(c *Deferrable_optionContext)

	// ExitConstraint_check_time is called when exiting the constraint_check_time production.
	ExitConstraint_check_time(c *Constraint_check_timeContext)

	// ExitTable_constraint_clause is called when exiting the table_constraint_clause production.
	ExitTable_constraint_clause(c *Table_constraint_clauseContext)

	// ExitTable_constraint is called when exiting the table_constraint production.
	ExitTable_constraint(c *Table_constraintContext)

	// ExitUsing_index_clause is called when exiting the using_index_clause production.
	ExitUsing_index_clause(c *Using_index_clauseContext)

	// ExitForeign_key_clause is called when exiting the foreign_key_clause production.
	ExitForeign_key_clause(c *Foreign_key_clauseContext)

	// ExitAlter_trigger_stmt is called when exiting the alter_trigger_stmt production.
	ExitAlter_trigger_stmt(c *Alter_trigger_stmtContext)

	// ExitAlter_trigger_option is called when exiting the alter_trigger_option production.
	ExitAlter_trigger_option(c *Alter_trigger_optionContext)

	// ExitAlter_table_partition_action_options is called when exiting the alter_table_partition_action_options production.
	ExitAlter_table_partition_action_options(c *Alter_table_partition_action_optionsContext)

	// ExitRefresh_materialized_view_stmt is called when exiting the refresh_materialized_view_stmt production.
	ExitRefresh_materialized_view_stmt(c *Refresh_materialized_view_stmtContext)

	// ExitComplete_del_null is called when exiting the complete_del_null production.
	ExitComplete_del_null(c *Complete_del_nullContext)

	// ExitRefresh_complete_del is called when exiting the refresh_complete_del production.
	ExitRefresh_complete_del(c *Refresh_complete_delContext)

	// ExitAlter_materialized_view_stmt is called when exiting the alter_materialized_view_stmt production.
	ExitAlter_materialized_view_stmt(c *Alter_materialized_view_stmtContext)

	// ExitAlter_view_stmt is called when exiting the alter_view_stmt production.
	ExitAlter_view_stmt(c *Alter_view_stmtContext)

	// ExitAlter_view_action is called when exiting the alter_view_action production.
	ExitAlter_view_action(c *Alter_view_actionContext)

	// ExitCons_novalidate is called when exiting the cons_novalidate production.
	ExitCons_novalidate(c *Cons_novalidateContext)

	// ExitView_constraint_clause is called when exiting the view_constraint_clause production.
	ExitView_constraint_clause(c *View_constraint_clauseContext)

	// ExitView_constraint is called when exiting the view_constraint production.
	ExitView_constraint(c *View_constraintContext)

	// ExitView_unique_spec is called when exiting the view_unique_spec production.
	ExitView_unique_spec(c *View_unique_specContext)

	// ExitView_refs_spec is called when exiting the view_refs_spec production.
	ExitView_refs_spec(c *View_refs_specContext)

	// ExitView_refs_spec_action is called when exiting the view_refs_spec_action production.
	ExitView_refs_spec_action(c *View_refs_spec_actionContext)

	// ExitCall_proc_stmt is called when exiting the call_proc_stmt production.
	ExitCall_proc_stmt(c *Call_proc_stmtContext)

	// ExitRaw_call_proc_stmt is called when exiting the raw_call_proc_stmt production.
	ExitRaw_call_proc_stmt(c *Raw_call_proc_stmtContext)

	// ExitCall_proc_stmt_2 is called when exiting the call_proc_stmt_2 production.
	ExitCall_proc_stmt_2(c *Call_proc_stmt_2Context)

	// ExitExec_proc_stmt is called when exiting the exec_proc_stmt production.
	ExitExec_proc_stmt(c *Exec_proc_stmtContext)

	// ExitDblink_clause is called when exiting the dblink_clause production.
	ExitDblink_clause(c *Dblink_clauseContext)

	// ExitDblink_clause2 is called when exiting the dblink_clause2 production.
	ExitDblink_clause2(c *Dblink_clause2Context)

	// ExitParam_list is called when exiting the param_list production.
	ExitParam_list(c *Param_listContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitRaw_exp_list is called when exiting the raw_exp_list production.
	ExitRaw_exp_list(c *Raw_exp_listContext)

	// ExitExp_list_2 is called when exiting the exp_list_2 production.
	ExitExp_list_2(c *Exp_list_2Context)

	// ExitExp_list is called when exiting the exp_list production.
	ExitExp_list(c *Exp_listContext)

	// ExitIns_exp_list is called when exiting the ins_exp_list production.
	ExitIns_exp_list(c *Ins_exp_listContext)

	// ExitLt_exp is called when exiting the lt_exp production.
	ExitLt_exp(c *Lt_expContext)

	// ExitRange_partition_exp is called when exiting the range_partition_exp production.
	ExitRange_partition_exp(c *Range_partition_expContext)

	// ExitRange_partition_exp_list is called when exiting the range_partition_exp_list production.
	ExitRange_partition_exp_list(c *Range_partition_exp_listContext)

	// ExitList_partition_exp is called when exiting the list_partition_exp production.
	ExitList_partition_exp(c *List_partition_expContext)

	// ExitList_partition_exp_list is called when exiting the list_partition_exp_list production.
	ExitList_partition_exp_list(c *List_partition_exp_listContext)

	// ExitList_partition_value_list is called when exiting the list_partition_value_list production.
	ExitList_partition_value_list(c *List_partition_value_listContext)

	// ExitClose_cursor_stmt is called when exiting the close_cursor_stmt production.
	ExitClose_cursor_stmt(c *Close_cursor_stmtContext)

	// ExitClose_cursor_statement is called when exiting the close_cursor_statement production.
	ExitClose_cursor_statement(c *Close_cursor_statementContext)

	// ExitBegin_trans_stmt is called when exiting the begin_trans_stmt production.
	ExitBegin_trans_stmt(c *Begin_trans_stmtContext)

	// ExitCommit_trans_stmt is called when exiting the commit_trans_stmt production.
	ExitCommit_trans_stmt(c *Commit_trans_stmtContext)

	// ExitCommit_head is called when exiting the commit_head production.
	ExitCommit_head(c *Commit_headContext)

	// ExitCommit_tail is called when exiting the commit_tail production.
	ExitCommit_tail(c *Commit_tailContext)

	// ExitCommit_wait_immed_option is called when exiting the commit_wait_immed_option production.
	ExitCommit_wait_immed_option(c *Commit_wait_immed_optionContext)

	// ExitConnect_stmt is called when exiting the connect_stmt production.
	ExitConnect_stmt(c *Connect_stmtContext)

	// ExitPassword is called when exiting the password production.
	ExitPassword(c *PasswordContext)

	// ExitTs_storage is called when exiting the ts_storage production.
	ExitTs_storage(c *Ts_storageContext)

	// ExitTs_storage_clause is called when exiting the ts_storage_clause production.
	ExitTs_storage_clause(c *Ts_storage_clauseContext)

	// ExitCreate_tablespace_stmt is called when exiting the create_tablespace_stmt production.
	ExitCreate_tablespace_stmt(c *Create_tablespace_stmtContext)

	// ExitCtss_with_clause is called when exiting the ctss_with_clause production.
	ExitCtss_with_clause(c *Ctss_with_clauseContext)

	// ExitCreate_tablespace_set_stmt is called when exiting the create_tablespace_set_stmt production.
	ExitCreate_tablespace_set_stmt(c *Create_tablespace_set_stmtContext)

	// ExitAlter_tablespace_set_stmt is called when exiting the alter_tablespace_set_stmt production.
	ExitAlter_tablespace_set_stmt(c *Alter_tablespace_set_stmtContext)

	// ExitCache is called when exiting the cache production.
	ExitCache(c *CacheContext)

	// ExitAlter_tablespace_stmt is called when exiting the alter_tablespace_stmt production.
	ExitAlter_tablespace_stmt(c *Alter_tablespace_stmtContext)

	// ExitKeep is called when exiting the keep production.
	ExitKeep(c *KeepContext)

	// ExitAlter_tablespace_action is called when exiting the alter_tablespace_action production.
	ExitAlter_tablespace_action(c *Alter_tablespace_actionContext)

	// ExitFile_list is called when exiting the file_list production.
	ExitFile_list(c *File_listContext)

	// ExitPathname_list is called when exiting the pathname_list production.
	ExitPathname_list(c *Pathname_listContext)

	// ExitInteger_list is called when exiting the integer_list production.
	ExitInteger_list(c *Integer_listContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitMirror is called when exiting the mirror production.
	ExitMirror(c *MirrorContext)

	// ExitAutoextend_nextsize is called when exiting the autoextend_nextsize production.
	ExitAutoextend_nextsize(c *Autoextend_nextsizeContext)

	// ExitAutoextend_maxsize is called when exiting the autoextend_maxsize production.
	ExitAutoextend_maxsize(c *Autoextend_maxsizeContext)

	// ExitAutoextend is called when exiting the autoextend production.
	ExitAutoextend(c *AutoextendContext)

	// ExitOn_raft is called when exiting the on_raft production.
	ExitOn_raft(c *On_raftContext)

	// ExitArchfile is called when exiting the archfile production.
	ExitArchfile(c *ArchfileContext)

	// ExitArchflag is called when exiting the archflag production.
	ExitArchflag(c *ArchflagContext)

	// ExitArchstyle_options is called when exiting the archstyle_options production.
	ExitArchstyle_options(c *Archstyle_optionsContext)

	// ExitArchstyle is called when exiting the archstyle production.
	ExitArchstyle(c *ArchstyleContext)

	// ExitArchdir is called when exiting the archdir production.
	ExitArchdir(c *ArchdirContext)

	// ExitBakfile is called when exiting the bakfile production.
	ExitBakfile(c *BakfileContext)

	// ExitParameters_option_list is called when exiting the parameters_option_list production.
	ExitParameters_option_list(c *Parameters_option_listContext)

	// ExitParameter_option_list is called when exiting the parameter_option_list production.
	ExitParameter_option_list(c *Parameter_option_listContext)

	// ExitParameter_option is called when exiting the parameter_option production.
	ExitParameter_option(c *Parameter_optionContext)

	// ExitPathname is called when exiting the pathname production.
	ExitPathname(c *PathnameContext)

	// ExitPathname_options is called when exiting the pathname_options production.
	ExitPathname_options(c *Pathname_optionsContext)

	// ExitBackup_stmt is called when exiting the backup_stmt production.
	ExitBackup_stmt(c *Backup_stmtContext)

	// ExitBack_range_option is called when exiting the back_range_option production.
	ExitBack_range_option(c *Back_range_optionContext)

	// ExitBack_archive_spec_null is called when exiting the back_archive_spec_null production.
	ExitBack_archive_spec_null(c *Back_archive_spec_nullContext)

	// ExitNot_backed_up is called when exiting the not_backed_up production.
	ExitNot_backed_up(c *Not_backed_upContext)

	// ExitArchive_spec is called when exiting the archive_spec production.
	ExitArchive_spec(c *Archive_specContext)

	// ExitSpec_lsn is called when exiting the spec_lsn production.
	ExitSpec_lsn(c *Spec_lsnContext)

	// ExitBackup_delete_archive is called when exiting the backup_delete_archive production.
	ExitBackup_delete_archive(c *Backup_delete_archiveContext)

	// ExitBackup_options is called when exiting the backup_options production.
	ExitBackup_options(c *Backup_optionsContext)

	// ExitCumulative is called when exiting the cumulative production.
	ExitCumulative(c *CumulativeContext)

	// ExitWith_bak_dir_list is called when exiting the with_bak_dir_list production.
	ExitWith_bak_dir_list(c *With_bak_dir_listContext)

	// ExitBase_on_backup is called when exiting the base_on_backup production.
	ExitBase_on_backup(c *Base_on_backupContext)

	// ExitBackup_to_options is called when exiting the backup_to_options production.
	ExitBackup_to_options(c *Backup_to_optionsContext)

	// ExitBackup_path_null is called when exiting the backup_path_null production.
	ExitBackup_path_null(c *Backup_path_nullContext)

	// ExitDevice_type is called when exiting the device_type production.
	ExitDevice_type(c *Device_typeContext)

	// ExitParms_command is called when exiting the parms_command production.
	ExitParms_command(c *Parms_commandContext)

	// ExitMedia_name is called when exiting the media_name production.
	ExitMedia_name(c *Media_nameContext)

	// ExitBackup_desc_options is called when exiting the backup_desc_options production.
	ExitBackup_desc_options(c *Backup_desc_optionsContext)

	// ExitBackup_desc is called when exiting the backup_desc production.
	ExitBackup_desc(c *Backup_descContext)

	// ExitBackup_maxsize is called when exiting the backup_maxsize production.
	ExitBackup_maxsize(c *Backup_maxsizeContext)

	// ExitBackup_limit is called when exiting the backup_limit production.
	ExitBackup_limit(c *Backup_limitContext)

	// ExitBackup_identified is called when exiting the backup_identified production.
	ExitBackup_identified(c *Backup_identifiedContext)

	// ExitBackup_compressed is called when exiting the backup_compressed production.
	ExitBackup_compressed(c *Backup_compressedContext)

	// ExitBackup_without is called when exiting the backup_without production.
	ExitBackup_without(c *Backup_withoutContext)

	// ExitBackup_tsk_thread_num_null is called when exiting the backup_tsk_thread_num_null production.
	ExitBackup_tsk_thread_num_null(c *Backup_tsk_thread_num_nullContext)

	// ExitBackup_parallel_dir is called when exiting the backup_parallel_dir production.
	ExitBackup_parallel_dir(c *Backup_parallel_dirContext)

	// ExitBackup_trace_file_level is called when exiting the backup_trace_file_level production.
	ExitBackup_trace_file_level(c *Backup_trace_file_levelContext)

	// ExitRestore_stmt is called when exiting the restore_stmt production.
	ExitRestore_stmt(c *Restore_stmtContext)

	// ExitRestore_datafile_lst is called when exiting the restore_datafile_lst production.
	ExitRestore_datafile_lst(c *Restore_datafile_lstContext)

	// ExitRestore_mapped_file is called when exiting the restore_mapped_file production.
	ExitRestore_mapped_file(c *Restore_mapped_fileContext)

	// ExitMapped_file is called when exiting the mapped_file production.
	ExitMapped_file(c *Mapped_fileContext)

	// ExitRes_struct is called when exiting the res_struct production.
	ExitRes_struct(c *Res_structContext)

	// ExitTsk_thread_num is called when exiting the tsk_thread_num production.
	ExitTsk_thread_num(c *Tsk_thread_numContext)

	// ExitRestore_tsk_thread_num_null is called when exiting the restore_tsk_thread_num_null production.
	ExitRestore_tsk_thread_num_null(c *Restore_tsk_thread_num_nullContext)

	// ExitRestore_parallel is called when exiting the restore_parallel production.
	ExitRestore_parallel(c *Restore_parallelContext)

	// ExitFull_table_name_options is called when exiting the full_table_name_options production.
	ExitFull_table_name_options(c *Full_table_name_optionsContext)

	// ExitRes_without_index_constraint is called when exiting the res_without_index_constraint production.
	ExitRes_without_index_constraint(c *Res_without_index_constraintContext)

	// ExitRestore_from is called when exiting the restore_from production.
	ExitRestore_from(c *Restore_fromContext)

	// ExitRes_until is called when exiting the res_until production.
	ExitRes_until(c *Res_untilContext)

	// ExitRestore_file_list_options is called when exiting the restore_file_list_options production.
	ExitRestore_file_list_options(c *Restore_file_list_optionsContext)

	// ExitMirror_file_list_options is called when exiting the mirror_file_list_options production.
	ExitMirror_file_list_options(c *Mirror_file_list_optionsContext)

	// ExitRestore_trace_file_level is called when exiting the restore_trace_file_level production.
	ExitRestore_trace_file_level(c *Restore_trace_file_levelContext)

	// ExitRestore_file_list is called when exiting the restore_file_list production.
	ExitRestore_file_list(c *Restore_file_listContext)

	// ExitRestore_file is called when exiting the restore_file production.
	ExitRestore_file(c *Restore_fileContext)

	// ExitMirror_file_list is called when exiting the mirror_file_list production.
	ExitMirror_file_list(c *Mirror_file_listContext)

	// ExitMirror_file is called when exiting the mirror_file production.
	ExitMirror_file(c *Mirror_fileContext)

	// ExitWith_bak_arch_dir_list is called when exiting the with_bak_arch_dir_list production.
	ExitWith_bak_arch_dir_list(c *With_bak_arch_dir_listContext)

	// ExitRestore_identified is called when exiting the restore_identified production.
	ExitRestore_identified(c *Restore_identifiedContext)

	// ExitCreate_func_stmt is called when exiting the create_func_stmt production.
	ExitCreate_func_stmt(c *Create_func_stmtContext)

	// ExitFunc_aggr_clause is called when exiting the func_aggr_clause production.
	ExitFunc_aggr_clause(c *Func_aggr_clauseContext)

	// ExitPipelined_options is called when exiting the pipelined_options production.
	ExitPipelined_options(c *Pipelined_optionsContext)

	// ExitReplace_option is called when exiting the replace_option production.
	ExitReplace_option(c *Replace_optionContext)

	// ExitEdit_options is called when exiting the edit_options production.
	ExitEdit_options(c *Edit_optionsContext)

	// ExitEncryption_option is called when exiting the encryption_option production.
	ExitEncryption_option(c *Encryption_optionContext)

	// ExitCalc_option is called when exiting the calc_option production.
	ExitCalc_option(c *Calc_optionContext)

	// ExitFunc_action is called when exiting the func_action production.
	ExitFunc_action(c *Func_actionContext)

	// ExitFunc_call_options is called when exiting the func_call_options production.
	ExitFunc_call_options(c *Func_call_optionsContext)

	// ExitFunc_call_option_list is called when exiting the func_call_option_list production.
	ExitFunc_call_option_list(c *Func_call_option_listContext)

	// ExitFunc_call_option is called when exiting the func_call_option production.
	ExitFunc_call_option(c *Func_call_optionContext)

	// ExitInvoker_rights_clause_options is called when exiting the invoker_rights_clause_options production.
	ExitInvoker_rights_clause_options(c *Invoker_rights_clause_optionsContext)

	// ExitInvoker_rights_clause is called when exiting the invoker_rights_clause production.
	ExitInvoker_rights_clause(c *Invoker_rights_clauseContext)

	// ExitDeterministic_clause_options is called when exiting the deterministic_clause_options production.
	ExitDeterministic_clause_options(c *Deterministic_clause_optionsContext)

	// ExitDeterministic_clause is called when exiting the deterministic_clause production.
	ExitDeterministic_clause(c *Deterministic_clauseContext)

	// ExitFunc_call_option2_options is called when exiting the func_call_option2_options production.
	ExitFunc_call_option2_options(c *Func_call_option2_optionsContext)

	// ExitFunc_call_option_list2 is called when exiting the func_call_option_list2 production.
	ExitFunc_call_option_list2(c *Func_call_option_list2Context)

	// ExitFunc_call_option2 is called when exiting the func_call_option2 production.
	ExitFunc_call_option2(c *Func_call_option2Context)

	// ExitResult_cache_clause is called when exiting the result_cache_clause production.
	ExitResult_cache_clause(c *Result_cache_clauseContext)

	// ExitInner_fun_name is called when exiting the inner_fun_name production.
	ExitInner_fun_name(c *Inner_fun_nameContext)

	// ExitPlatform_type is called when exiting the platform_type production.
	ExitPlatform_type(c *Platform_typeContext)

	// ExitParam_def_list_option is called when exiting the param_def_list_option production.
	ExitParam_def_list_option(c *Param_def_list_optionContext)

	// ExitParam_def_list is called when exiting the param_def_list production.
	ExitParam_def_list(c *Param_def_listContext)

	// ExitParam_def is called when exiting the param_def production.
	ExitParam_def(c *Param_defContext)

	// ExitParam_in_out_option is called when exiting the param_in_out_option production.
	ExitParam_in_out_option(c *Param_in_out_optionContext)

	// ExitIs_as is called when exiting the is_as production.
	ExitIs_as(c *Is_asContext)

	// ExitStat_on_org_stmt is called when exiting the stat_on_org_stmt production.
	ExitStat_on_org_stmt(c *Stat_on_org_stmtContext)

	// ExitStat_size is called when exiting the stat_size production.
	ExitStat_size(c *Stat_sizeContext)

	// ExitStat_para is called when exiting the stat_para production.
	ExitStat_para(c *Stat_paraContext)

	// ExitStat_summarize is called when exiting the stat_summarize production.
	ExitStat_summarize(c *Stat_summarizeContext)

	// ExitMstat_ex is called when exiting the mstat_ex production.
	ExitMstat_ex(c *Mstat_exContext)

	// ExitIndexid is called when exiting the indexid production.
	ExitIndexid(c *IndexidContext)

	// ExitGlobal_tag is called when exiting the global_tag production.
	ExitGlobal_tag(c *Global_tagContext)

	// ExitBm_join_index_clause is called when exiting the bm_join_index_clause production.
	ExitBm_join_index_clause(c *Bm_join_index_clauseContext)

	// ExitParallel_stmt is called when exiting the parallel_stmt production.
	ExitParallel_stmt(c *Parallel_stmtContext)

	// ExitCreate_index_stmt is called when exiting the create_index_stmt production.
	ExitCreate_index_stmt(c *Create_index_stmtContext)

	// ExitWith_inner is called when exiting the with_inner production.
	ExitWith_inner(c *With_innerContext)

	// ExitIndex_no_sort is called when exiting the index_no_sort production.
	ExitIndex_no_sort(c *Index_no_sortContext)

	// ExitOnline_options is called when exiting the online_options production.
	ExitOnline_options(c *Online_optionsContext)

	// ExitUnusable_options is called when exiting the unusable_options production.
	ExitUnusable_options(c *Unusable_optionsContext)

	// ExitReverse_options is called when exiting the reverse_options production.
	ExitReverse_options(c *Reverse_optionsContext)

	// ExitIndex_column_list is called when exiting the index_column_list production.
	ExitIndex_column_list(c *Index_column_listContext)

	// ExitIndex_column_name is called when exiting the index_column_name production.
	ExitIndex_column_name(c *Index_column_nameContext)

	// ExitStorage_hash_tag is called when exiting the storage_hash_tag production.
	ExitStorage_hash_tag(c *Storage_hash_tagContext)

	// ExitStorage_hash is called when exiting the storage_hash production.
	ExitStorage_hash(c *Storage_hashContext)

	// ExitStorage_tag is called when exiting the storage_tag production.
	ExitStorage_tag(c *Storage_tagContext)

	// ExitStorage_tag_nn is called when exiting the storage_tag_nn production.
	ExitStorage_tag_nn(c *Storage_tag_nnContext)

	// ExitTablespace_clause is called when exiting the tablespace_clause production.
	ExitTablespace_clause(c *Tablespace_clauseContext)

	// ExitObject_table_substitution_clause is called when exiting the object_table_substitution_clause production.
	ExitObject_table_substitution_clause(c *Object_table_substitution_clauseContext)

	// ExitObject_table_substitution is called when exiting the object_table_substitution production.
	ExitObject_table_substitution(c *Object_table_substitutionContext)

	// ExitOid_clause is called when exiting the oid_clause production.
	ExitOid_clause(c *Oid_clauseContext)

	// ExitOid_gen_type is called when exiting the oid_gen_type production.
	ExitOid_gen_type(c *Oid_gen_typeContext)

	// ExitOid_index_clause is called when exiting the oid_index_clause production.
	ExitOid_index_clause(c *Oid_index_clauseContext)

	// ExitOid_tablespace_clause is called when exiting the oid_tablespace_clause production.
	ExitOid_tablespace_clause(c *Oid_tablespace_clauseContext)

	// ExitOid_tablespace_name is called when exiting the oid_tablespace_name production.
	ExitOid_tablespace_name(c *Oid_tablespace_nameContext)

	// ExitLocal_option is called when exiting the local_option production.
	ExitLocal_option(c *Local_optionContext)

	// ExitStorage_list is called when exiting the storage_list production.
	ExitStorage_list(c *Storage_listContext)

	// ExitStorage_hashpartmap is called when exiting the storage_hashpartmap production.
	ExitStorage_hashpartmap(c *Storage_hashpartmapContext)

	// ExitStorage is called when exiting the storage production.
	ExitStorage(c *StorageContext)

	// ExitId_list is called when exiting the id_list production.
	ExitId_list(c *Id_listContext)

	// ExitCreate_proc_stmt is called when exiting the create_proc_stmt production.
	ExitCreate_proc_stmt(c *Create_proc_stmtContext)

	// ExitCreate_package_stmt is called when exiting the create_package_stmt production.
	ExitCreate_package_stmt(c *Create_package_stmtContext)

	// ExitPkg_cls_flag is called when exiting the pkg_cls_flag production.
	ExitPkg_cls_flag(c *Pkg_cls_flagContext)

	// ExitBlk_end_option is called when exiting the blk_end_option production.
	ExitBlk_end_option(c *Blk_end_optionContext)

	// ExitBlk_end_option_low is called when exiting the blk_end_option_low production.
	ExitBlk_end_option_low(c *Blk_end_option_lowContext)

	// ExitPackage_def_list_options is called when exiting the package_def_list_options production.
	ExitPackage_def_list_options(c *Package_def_list_optionsContext)

	// ExitPackage_def_list is called when exiting the package_def_list production.
	ExitPackage_def_list(c *Package_def_listContext)

	// ExitPackage_def is called when exiting the package_def production.
	ExitPackage_def(c *Package_defContext)

	// ExitRestrict_param_lst is called when exiting the restrict_param_lst production.
	ExitRestrict_param_lst(c *Restrict_param_lstContext)

	// ExitCreate_package_body_stmt is called when exiting the create_package_body_stmt production.
	ExitCreate_package_body_stmt(c *Create_package_body_stmtContext)

	// ExitCreate_pkg_body_header is called when exiting the create_pkg_body_header production.
	ExitCreate_pkg_body_header(c *Create_pkg_body_headerContext)

	// ExitPkg_cls_body_flag is called when exiting the pkg_cls_body_flag production.
	ExitPkg_cls_body_flag(c *Pkg_cls_body_flagContext)

	// ExitPackage_body_init_option is called when exiting the package_body_init_option production.
	ExitPackage_body_init_option(c *Package_body_init_optionContext)

	// ExitPackage_body_def_list is called when exiting the package_body_def_list production.
	ExitPackage_body_def_list(c *Package_body_def_listContext)

	// ExitPackage_body_def is called when exiting the package_body_def production.
	ExitPackage_body_def(c *Package_body_defContext)

	// ExitPackage_body_def2 is called when exiting the package_body_def2 production.
	ExitPackage_body_def2(c *Package_body_def2Context)

	// ExitCheck_exec_options is called when exiting the check_exec_options production.
	ExitCheck_exec_options(c *Check_exec_optionsContext)

	// ExitSubpg_decl_stmt is called when exiting the subpg_decl_stmt production.
	ExitSubpg_decl_stmt(c *Subpg_decl_stmtContext)

	// ExitCreate_type_stmt is called when exiting the create_type_stmt production.
	ExitCreate_type_stmt(c *Create_type_stmtContext)

	// ExitForce_option is called when exiting the force_option production.
	ExitForce_option(c *Force_optionContext)

	// ExitObject_option is called when exiting the object_option production.
	ExitObject_option(c *Object_optionContext)

	// ExitUnder_option is called when exiting the under_option production.
	ExitUnder_option(c *Under_optionContext)

	// ExitObject_def_list_options is called when exiting the object_def_list_options production.
	ExitObject_def_list_options(c *Object_def_list_optionsContext)

	// ExitObject_def_list is called when exiting the object_def_list production.
	ExitObject_def_list(c *Object_def_listContext)

	// ExitObject_def is called when exiting the object_def production.
	ExitObject_def(c *Object_defContext)

	// ExitMember_static is called when exiting the member_static production.
	ExitMember_static(c *Member_staticContext)

	// ExitMethod_inherit_options is called when exiting the method_inherit_options production.
	ExitMethod_inherit_options(c *Method_inherit_optionsContext)

	// ExitMethod_inherit_option is called when exiting the method_inherit_option production.
	ExitMethod_inherit_option(c *Method_inherit_optionContext)

	// ExitFinal_inst_list_options is called when exiting the final_inst_list_options production.
	ExitFinal_inst_list_options(c *Final_inst_list_optionsContext)

	// ExitFinal_inst_list is called when exiting the final_inst_list production.
	ExitFinal_inst_list(c *Final_inst_listContext)

	// ExitFinal_inst is called when exiting the final_inst production.
	ExitFinal_inst(c *Final_instContext)

	// ExitOverriding_option is called when exiting the overriding_option production.
	ExitOverriding_option(c *Overriding_optionContext)

	// ExitMethod_attr_options is called when exiting the method_attr_options production.
	ExitMethod_attr_options(c *Method_attr_optionsContext)

	// ExitMethod_attr is called when exiting the method_attr production.
	ExitMethod_attr(c *Method_attrContext)

	// ExitCreate_type_body_stmt is called when exiting the create_type_body_stmt production.
	ExitCreate_type_body_stmt(c *Create_type_body_stmtContext)

	// ExitObject_body_def_list is called when exiting the object_body_def_list production.
	ExitObject_body_def_list(c *Object_body_def_listContext)

	// ExitObject_body_def is called when exiting the object_body_def production.
	ExitObject_body_def(c *Object_body_defContext)

	// ExitCreate_context_stmt is called when exiting the create_context_stmt production.
	ExitCreate_context_stmt(c *Create_context_stmtContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitInitialized is called when exiting the initialized production.
	ExitInitialized(c *InitializedContext)

	// ExitCreate_directory_stmt is called when exiting the create_directory_stmt production.
	ExitCreate_directory_stmt(c *Create_directory_stmtContext)

	// ExitCreate_crypto_stmt is called when exiting the create_crypto_stmt production.
	ExitCreate_crypto_stmt(c *Create_crypto_stmtContext)

	// ExitAlter_crypto_stmt is called when exiting the alter_crypto_stmt production.
	ExitAlter_crypto_stmt(c *Alter_crypto_stmtContext)

	// ExitAlter_crypto_action is called when exiting the alter_crypto_action production.
	ExitAlter_crypto_action(c *Alter_crypto_actionContext)

	// ExitComment_stmt is called when exiting the comment_stmt production.
	ExitComment_stmt(c *Comment_stmtContext)

	// ExitComment_tag is called when exiting the comment_tag production.
	ExitComment_tag(c *Comment_tagContext)

	// ExitCreate_partition_group_stmt is called when exiting the create_partition_group_stmt production.
	ExitCreate_partition_group_stmt(c *Create_partition_group_stmtContext)

	// ExitStorage_act_datatype is called when exiting the storage_act_datatype production.
	ExitStorage_act_datatype(c *Storage_act_datatypeContext)

	// ExitPg_storage_lst is called when exiting the pg_storage_lst production.
	ExitPg_storage_lst(c *Pg_storage_lstContext)

	// ExitPg_storage is called when exiting the pg_storage production.
	ExitPg_storage(c *Pg_storageContext)

	// ExitCreate_table_stmt is called when exiting the create_table_stmt production.
	ExitCreate_table_stmt(c *Create_table_stmtContext)

	// ExitCtab_append_attr_clause is called when exiting the ctab_append_attr_clause production.
	ExitCtab_append_attr_clause(c *Ctab_append_attr_clauseContext)

	// ExitCtab_append_attr_list is called when exiting the ctab_append_attr_list production.
	ExitCtab_append_attr_list(c *Ctab_append_attr_listContext)

	// ExitCobjtab_append_attr_clause is called when exiting the cobjtab_append_attr_clause production.
	ExitCobjtab_append_attr_clause(c *Cobjtab_append_attr_clauseContext)

	// ExitCobjtab_append_attr_list is called when exiting the cobjtab_append_attr_list production.
	ExitCobjtab_append_attr_list(c *Cobjtab_append_attr_listContext)

	// ExitCtab_append_attr is called when exiting the ctab_append_attr production.
	ExitCtab_append_attr(c *Ctab_append_attrContext)

	// ExitCobjtab_append_attr is called when exiting the cobjtab_append_attr production.
	ExitCobjtab_append_attr(c *Cobjtab_append_attrContext)

	// ExitCreate_table_action is called when exiting the create_table_action production.
	ExitCreate_table_action(c *Create_table_actionContext)

	// ExitCtab_log_options is called when exiting the ctab_log_options production.
	ExitCtab_log_options(c *Ctab_log_optionsContext)

	// ExitCtab_log_option is called when exiting the ctab_log_option production.
	ExitCtab_log_option(c *Ctab_log_optionContext)

	// ExitCtab_error_options is called when exiting the ctab_error_options production.
	ExitCtab_error_options(c *Ctab_error_optionsContext)

	// ExitAdvance_log_clause is called when exiting the advance_log_clause production.
	ExitAdvance_log_clause(c *Advance_log_clauseContext)

	// ExitAdd_log_clause is called when exiting the add_log_clause production.
	ExitAdd_log_clause(c *Add_log_clauseContext)

	// ExitCtab_error_option is called when exiting the ctab_error_option production.
	ExitCtab_error_option(c *Ctab_error_optionContext)

	// ExitCtab_row_movement_clause is called when exiting the ctab_row_movement_clause production.
	ExitCtab_row_movement_clause(c *Ctab_row_movement_clauseContext)

	// ExitRange_distribute_act is called when exiting the range_distribute_act production.
	ExitRange_distribute_act(c *Range_distribute_actContext)

	// ExitRange_distribute_act_lst is called when exiting the range_distribute_act_lst production.
	ExitRange_distribute_act_lst(c *Range_distribute_act_lstContext)

	// ExitList_distribute_act is called when exiting the list_distribute_act production.
	ExitList_distribute_act(c *List_distribute_actContext)

	// ExitList_distribute_act_list is called when exiting the list_distribute_act_list production.
	ExitList_distribute_act_list(c *List_distribute_act_listContext)

	// ExitDistribute_by_option is called when exiting the distribute_by_option production.
	ExitDistribute_by_option(c *Distribute_by_optionContext)

	// ExitDistribute_by is called when exiting the distribute_by production.
	ExitDistribute_by(c *Distribute_byContext)

	// ExitIncrement_set is called when exiting the increment_set production.
	ExitIncrement_set(c *Increment_setContext)

	// ExitIncrement is called when exiting the increment production.
	ExitIncrement(c *IncrementContext)

	// ExitRowdependencies_clause is called when exiting the rowdependencies_clause production.
	ExitRowdependencies_clause(c *Rowdependencies_clauseContext)

	// ExitPg_sub_partition is called when exiting the pg_sub_partition production.
	ExitPg_sub_partition(c *Pg_sub_partitionContext)

	// ExitTable_type_option is called when exiting the table_type_option production.
	ExitTable_type_option(c *Table_type_optionContext)

	// ExitTable_temp_option is called when exiting the table_temp_option production.
	ExitTable_temp_option(c *Table_temp_optionContext)

	// ExitObjtab_elem_constraint is called when exiting the objtab_elem_constraint production.
	ExitObjtab_elem_constraint(c *Objtab_elem_constraintContext)

	// ExitObjtab_element_cons_list is called when exiting the objtab_element_cons_list production.
	ExitObjtab_element_cons_list(c *Objtab_element_cons_listContext)

	// ExitObjtab_element_cons is called when exiting the objtab_element_cons production.
	ExitObjtab_element_cons(c *Objtab_element_consContext)

	// ExitObjcol_constraint is called when exiting the objcol_constraint production.
	ExitObjcol_constraint(c *Objcol_constraintContext)

	// ExitTable_element_list is called when exiting the table_element_list production.
	ExitTable_element_list(c *Table_element_listContext)

	// ExitTable_element is called when exiting the table_element production.
	ExitTable_element(c *Table_elementContext)

	// ExitTable_constraint_def is called when exiting the table_constraint_def production.
	ExitTable_constraint_def(c *Table_constraint_defContext)

	// ExitOn_commit_option is called when exiting the on_commit_option production.
	ExitOn_commit_option(c *On_commit_optionContext)

	// ExitOn_commit_option_nn is called when exiting the on_commit_option_nn production.
	ExitOn_commit_option_nn(c *On_commit_option_nnContext)

	// ExitLogging_option is called when exiting the logging_option production.
	ExitLogging_option(c *Logging_optionContext)

	// ExitLogging_option_nn is called when exiting the logging_option_nn production.
	ExitLogging_option_nn(c *Logging_option_nnContext)

	// ExitPartition_clause is called when exiting the partition_clause production.
	ExitPartition_clause(c *Partition_clauseContext)

	// ExitPartition_clause_nn is called when exiting the partition_clause_nn production.
	ExitPartition_clause_nn(c *Partition_clause_nnContext)

	// ExitHorizon_partition_clause is called when exiting the horizon_partition_clause production.
	ExitHorizon_partition_clause(c *Horizon_partition_clauseContext)

	// ExitCompress_tag_hdr is called when exiting the compress_tag_hdr production.
	ExitCompress_tag_hdr(c *Compress_tag_hdrContext)

	// ExitCompress_clause_opt is called when exiting the compress_clause_opt production.
	ExitCompress_clause_opt(c *Compress_clause_optContext)

	// ExitCompress_tag is called when exiting the compress_tag production.
	ExitCompress_tag(c *Compress_tagContext)

	// ExitCompress_level is called when exiting the compress_level production.
	ExitCompress_level(c *Compress_levelContext)

	// ExitCompress_type is called when exiting the compress_type production.
	ExitCompress_type(c *Compress_typeContext)

	// ExitRange_partition is called when exiting the range_partition production.
	ExitRange_partition(c *Range_partitionContext)

	// ExitRange_partition_list is called when exiting the range_partition_list production.
	ExitRange_partition_list(c *Range_partition_listContext)

	// ExitHash_partition is called when exiting the hash_partition production.
	ExitHash_partition(c *Hash_partitionContext)

	// ExitHash_partition_list is called when exiting the hash_partition_list production.
	ExitHash_partition_list(c *Hash_partition_listContext)

	// ExitList_partition is called when exiting the list_partition production.
	ExitList_partition(c *List_partitionContext)

	// ExitList_partition_list is called when exiting the list_partition_list production.
	ExitList_partition_list(c *List_partition_listContext)

	// ExitSplit_partition_list is called when exiting the split_partition_list production.
	ExitSplit_partition_list(c *Split_partition_listContext)

	// ExitPartition_act is called when exiting the partition_act production.
	ExitPartition_act(c *Partition_actContext)

	// ExitVertical_partition_act is called when exiting the vertical_partition_act production.
	ExitVertical_partition_act(c *Vertical_partition_actContext)

	// ExitInterval_item is called when exiting the interval_item production.
	ExitInterval_item(c *Interval_itemContext)

	// ExitHorizon_partition_act_datatype is called when exiting the horizon_partition_act_datatype production.
	ExitHorizon_partition_act_datatype(c *Horizon_partition_act_datatypeContext)

	// ExitHorizon_partition_act is called when exiting the horizon_partition_act production.
	ExitHorizon_partition_act(c *Horizon_partition_actContext)

	// ExitLock_partitions_clause is called when exiting the lock_partitions_clause production.
	ExitLock_partitions_clause(c *Lock_partitions_clauseContext)

	// ExitSubpartion_template_list_datatype_options is called when exiting the subpartion_template_list_datatype_options production.
	ExitSubpartion_template_list_datatype_options(c *Subpartion_template_list_datatype_optionsContext)

	// ExitSubpartion_template_list_datatype is called when exiting the subpartion_template_list_datatype production.
	ExitSubpartion_template_list_datatype(c *Subpartion_template_list_datatypeContext)

	// ExitSubpartion_template_list_options is called when exiting the subpartion_template_list_options production.
	ExitSubpartion_template_list_options(c *Subpartion_template_list_optionsContext)

	// ExitSubpartion_template_list is called when exiting the subpartion_template_list production.
	ExitSubpartion_template_list(c *Subpartion_template_listContext)

	// ExitSubpartion_template_datatype is called when exiting the subpartion_template_datatype production.
	ExitSubpartion_template_datatype(c *Subpartion_template_datatypeContext)

	// ExitRange_subpartion_template_datatype is called when exiting the range_subpartion_template_datatype production.
	ExitRange_subpartion_template_datatype(c *Range_subpartion_template_datatypeContext)

	// ExitHash_subpartion_template_datatype is called when exiting the hash_subpartion_template_datatype production.
	ExitHash_subpartion_template_datatype(c *Hash_subpartion_template_datatypeContext)

	// ExitHash_subpartions_template_datatype_option is called when exiting the hash_subpartions_template_datatype_option production.
	ExitHash_subpartions_template_datatype_option(c *Hash_subpartions_template_datatype_optionContext)

	// ExitList_subpartion_template_datatype is called when exiting the list_subpartion_template_datatype production.
	ExitList_subpartion_template_datatype(c *List_subpartion_template_datatypeContext)

	// ExitSubpartion_template is called when exiting the subpartion_template production.
	ExitSubpartion_template(c *Subpartion_templateContext)

	// ExitRange_subpartion_template is called when exiting the range_subpartion_template production.
	ExitRange_subpartion_template(c *Range_subpartion_templateContext)

	// ExitHash_subpartion_template is called when exiting the hash_subpartion_template production.
	ExitHash_subpartion_template(c *Hash_subpartion_templateContext)

	// ExitHash_subpartions_template_option is called when exiting the hash_subpartions_template_option production.
	ExitHash_subpartions_template_option(c *Hash_subpartions_template_optionContext)

	// ExitList_subpartion_template is called when exiting the list_subpartion_template production.
	ExitList_subpartion_template(c *List_subpartion_templateContext)

	// ExitRange_subpartition is called when exiting the range_subpartition production.
	ExitRange_subpartition(c *Range_subpartitionContext)

	// ExitHash_subpartition is called when exiting the hash_subpartition production.
	ExitHash_subpartition(c *Hash_subpartitionContext)

	// ExitList_subpartition is called when exiting the list_subpartition production.
	ExitList_subpartition(c *List_subpartitionContext)

	// ExitRange_subpartition_list is called when exiting the range_subpartition_list production.
	ExitRange_subpartition_list(c *Range_subpartition_listContext)

	// ExitHash_subpartition_list is called when exiting the hash_subpartition_list production.
	ExitHash_subpartition_list(c *Hash_subpartition_listContext)

	// ExitList_subpartition_list is called when exiting the list_subpartition_list production.
	ExitList_subpartition_list(c *List_subpartition_listContext)

	// ExitSubpartition_hash_desc is called when exiting the subpartition_hash_desc production.
	ExitSubpartition_hash_desc(c *Subpartition_hash_descContext)

	// ExitSubpartition_range_desc is called when exiting the subpartition_range_desc production.
	ExitSubpartition_range_desc(c *Subpartition_range_descContext)

	// ExitSubpartition_list_desc is called when exiting the subpartition_list_desc production.
	ExitSubpartition_list_desc(c *Subpartition_list_descContext)

	// ExitSubpartition_hash_desc_list is called when exiting the subpartition_hash_desc_list production.
	ExitSubpartition_hash_desc_list(c *Subpartition_hash_desc_listContext)

	// ExitSubpartition_range_desc_list is called when exiting the subpartition_range_desc_list production.
	ExitSubpartition_range_desc_list(c *Subpartition_range_desc_listContext)

	// ExitSubpartition_list_desc_list is called when exiting the subpartition_list_desc_list production.
	ExitSubpartition_list_desc_list(c *Subpartition_list_desc_listContext)

	// ExitSubpartition_desc is called when exiting the subpartition_desc production.
	ExitSubpartition_desc(c *Subpartition_descContext)

	// ExitSubpartition_desc_option is called when exiting the subpartition_desc_option production.
	ExitSubpartition_desc_option(c *Subpartition_desc_optionContext)

	// ExitAdd_subpartition_desc is called when exiting the add_subpartition_desc production.
	ExitAdd_subpartition_desc(c *Add_subpartition_descContext)

	// ExitPartition_no is called when exiting the partition_no production.
	ExitPartition_no(c *Partition_noContext)

	// ExitComment_clause is called when exiting the comment_clause production.
	ExitComment_clause(c *Comment_clauseContext)

	// ExitEncrypt_clause_options is called when exiting the encrypt_clause_options production.
	ExitEncrypt_clause_options(c *Encrypt_clause_optionsContext)

	// ExitEncrypt_clause is called when exiting the encrypt_clause production.
	ExitEncrypt_clause(c *Encrypt_clauseContext)

	// ExitEncrypt_cipher is called when exiting the encrypt_cipher production.
	ExitEncrypt_cipher(c *Encrypt_cipherContext)

	// ExitCrypto_name is called when exiting the crypto_name production.
	ExitCrypto_name(c *Crypto_nameContext)

	// ExitCipher_name is called when exiting the cipher_name production.
	ExitCipher_name(c *Cipher_nameContext)

	// ExitFull_cipher_name is called when exiting the full_cipher_name production.
	ExitFull_cipher_name(c *Full_cipher_nameContext)

	// ExitEncrypt_type is called when exiting the encrypt_type production.
	ExitEncrypt_type(c *Encrypt_typeContext)

	// ExitManual_clause is called when exiting the manual_clause production.
	ExitManual_clause(c *Manual_clauseContext)

	// ExitUser_clause_options is called when exiting the user_clause_options production.
	ExitUser_clause_options(c *User_clause_optionsContext)

	// ExitUser_clause is called when exiting the user_clause production.
	ExitUser_clause(c *User_clauseContext)

	// ExitUser_list_option is called when exiting the user_list_option production.
	ExitUser_list_option(c *User_list_optionContext)

	// ExitUser_list is called when exiting the user_list production.
	ExitUser_list(c *User_listContext)

	// ExitHash_cipher is called when exiting the hash_cipher production.
	ExitHash_cipher(c *Hash_cipherContext)

	// ExitHash_type is called when exiting the hash_type production.
	ExitHash_type(c *Hash_typeContext)

	// ExitSpace_limit is called when exiting the space_limit production.
	ExitSpace_limit(c *Space_limitContext)

	// ExitSpace_limit_nn is called when exiting the space_limit_nn production.
	ExitSpace_limit_nn(c *Space_limit_nnContext)

	// ExitSpace_limit_1 is called when exiting the space_limit_1 production.
	ExitSpace_limit_1(c *Space_limit_1Context)

	// ExitSpace_limit2 is called when exiting the space_limit2 production.
	ExitSpace_limit2(c *Space_limit2Context)

	// ExitDel_res is called when exiting the del_res production.
	ExitDel_res(c *Del_resContext)

	// ExitTrig_enable is called when exiting the trig_enable production.
	ExitTrig_enable(c *Trig_enableContext)

	// ExitAt_raft is called when exiting the at_raft production.
	ExitAt_raft(c *At_raftContext)

	// ExitCreate_trigger_stmt is called when exiting the create_trigger_stmt production.
	ExitCreate_trigger_stmt(c *Create_trigger_stmtContext)

	// ExitBefore_after is called when exiting the before_after production.
	ExitBefore_after(c *Before_afterContext)

	// ExitTrig_del_ins_upd_list is called when exiting the trig_del_ins_upd_list production.
	ExitTrig_del_ins_upd_list(c *Trig_del_ins_upd_listContext)

	// ExitTrig_del_ins_upd is called when exiting the trig_del_ins_upd production.
	ExitTrig_del_ins_upd(c *Trig_del_ins_updContext)

	// ExitUpdate_of_option is called when exiting the update_of_option production.
	ExitUpdate_of_option(c *Update_of_optionContext)

	// ExitNowait is called when exiting the nowait production.
	ExitNowait(c *NowaitContext)

	// ExitTrig_event_list is called when exiting the trig_event_list production.
	ExitTrig_event_list(c *Trig_event_listContext)

	// ExitTrig_event is called when exiting the trig_event production.
	ExitTrig_event(c *Trig_eventContext)

	// ExitEvent_object is called when exiting the event_object production.
	ExitEvent_object(c *Event_objectContext)

	// ExitTrig_referencing_def_options is called when exiting the trig_referencing_def_options production.
	ExitTrig_referencing_def_options(c *Trig_referencing_def_optionsContext)

	// ExitTrig_referencing_def is called when exiting the trig_referencing_def production.
	ExitTrig_referencing_def(c *Trig_referencing_defContext)

	// ExitTrig_referencing_list is called when exiting the trig_referencing_list production.
	ExitTrig_referencing_list(c *Trig_referencing_listContext)

	// ExitTrig_referencing_old is called when exiting the trig_referencing_old production.
	ExitTrig_referencing_old(c *Trig_referencing_oldContext)

	// ExitTrig_referencing_new is called when exiting the trig_referencing_new production.
	ExitTrig_referencing_new(c *Trig_referencing_newContext)

	// ExitTrig_for_each_option is called when exiting the trig_for_each_option production.
	ExitTrig_for_each_option(c *Trig_for_each_optionContext)

	// ExitTrig_timer_rate is called when exiting the trig_timer_rate production.
	ExitTrig_timer_rate(c *Trig_timer_rateContext)

	// ExitExec_ep_seqno is called when exiting the exec_ep_seqno production.
	ExitExec_ep_seqno(c *Exec_ep_seqnoContext)

	// ExitRate_over_day is called when exiting the rate_over_day production.
	ExitRate_over_day(c *Rate_over_dayContext)

	// ExitMonth_rate is called when exiting the month_rate production.
	ExitMonth_rate(c *Month_rateContext)

	// ExitDay_in_month is called when exiting the day_in_month production.
	ExitDay_in_month(c *Day_in_monthContext)

	// ExitDay_in_month_week is called when exiting the day_in_month_week production.
	ExitDay_in_month_week(c *Day_in_month_weekContext)

	// ExitWeek_rate is called when exiting the week_rate production.
	ExitWeek_rate(c *Week_rateContext)

	// ExitDay_of_week_list is called when exiting the day_of_week_list production.
	ExitDay_of_week_list(c *Day_of_week_listContext)

	// ExitDay_rate is called when exiting the day_rate production.
	ExitDay_rate(c *Day_rateContext)

	// ExitRate_in_day is called when exiting the rate_in_day production.
	ExitRate_in_day(c *Rate_in_dayContext)

	// ExitOnce_in_day is called when exiting the once_in_day production.
	ExitOnce_in_day(c *Once_in_dayContext)

	// ExitTimes_in_day is called when exiting the times_in_day production.
	ExitTimes_in_day(c *Times_in_dayContext)

	// ExitDuaring_time is called when exiting the duaring_time production.
	ExitDuaring_time(c *Duaring_timeContext)

	// ExitDuaring_date is called when exiting the duaring_date production.
	ExitDuaring_date(c *Duaring_dateContext)

	// ExitTrig_when_option is called when exiting the trig_when_option production.
	ExitTrig_when_option(c *Trig_when_optionContext)

	// ExitTrig_when_condition is called when exiting the trig_when_condition production.
	ExitTrig_when_condition(c *Trig_when_conditionContext)

	// ExitRepeat_interval_stmt is called when exiting the repeat_interval_stmt production.
	ExitRepeat_interval_stmt(c *Repeat_interval_stmtContext)

	// ExitMax_run_duration is called when exiting the max_run_duration production.
	ExitMax_run_duration(c *Max_run_durationContext)

	// ExitRepeat_interval is called when exiting the repeat_interval production.
	ExitRepeat_interval(c *Repeat_intervalContext)

	// ExitFrequency_clause is called when exiting the frequency_clause production.
	ExitFrequency_clause(c *Frequency_clauseContext)

	// ExitFrequency_define is called when exiting the frequency_define production.
	ExitFrequency_define(c *Frequency_defineContext)

	// ExitPredefined_frequency is called when exiting the predefined_frequency production.
	ExitPredefined_frequency(c *Predefined_frequencyContext)

	// ExitInterval_clause_list is called when exiting the interval_clause_list production.
	ExitInterval_clause_list(c *Interval_clause_listContext)

	// ExitInterval_clause_single is called when exiting the interval_clause_single production.
	ExitInterval_clause_single(c *Interval_clause_singleContext)

	// ExitInterval_clause is called when exiting the interval_clause production.
	ExitInterval_clause(c *Interval_clauseContext)

	// ExitIntervalnum is called when exiting the intervalnum production.
	ExitIntervalnum(c *IntervalnumContext)

	// ExitBymonth_clause is called when exiting the bymonth_clause production.
	ExitBymonth_clause(c *Bymonth_clauseContext)

	// ExitMonthlist is called when exiting the monthlist production.
	ExitMonthlist(c *MonthlistContext)

	// ExitMonth is called when exiting the month production.
	ExitMonth(c *MonthContext)

	// ExitNumeric_month is called when exiting the numeric_month production.
	ExitNumeric_month(c *Numeric_monthContext)

	// ExitChar_month is called when exiting the char_month production.
	ExitChar_month(c *Char_monthContext)

	// ExitByweekno_clause is called when exiting the byweekno_clause production.
	ExitByweekno_clause(c *Byweekno_clauseContext)

	// ExitWeekno_list is called when exiting the weekno_list production.
	ExitWeekno_list(c *Weekno_listContext)

	// ExitWeekno is called when exiting the weekno production.
	ExitWeekno(c *WeeknoContext)

	// ExitByyearday_clause is called when exiting the byyearday_clause production.
	ExitByyearday_clause(c *Byyearday_clauseContext)

	// ExitYearday_list is called when exiting the yearday_list production.
	ExitYearday_list(c *Yearday_listContext)

	// ExitYearday is called when exiting the yearday production.
	ExitYearday(c *YeardayContext)

	// ExitBymonthday_clause is called when exiting the bymonthday_clause production.
	ExitBymonthday_clause(c *Bymonthday_clauseContext)

	// ExitMonthday_list is called when exiting the monthday_list production.
	ExitMonthday_list(c *Monthday_listContext)

	// ExitMonthday is called when exiting the monthday production.
	ExitMonthday(c *MonthdayContext)

	// ExitByday_clause is called when exiting the byday_clause production.
	ExitByday_clause(c *Byday_clauseContext)

	// ExitByday_list is called when exiting the byday_list production.
	ExitByday_list(c *Byday_listContext)

	// ExitByday is called when exiting the byday production.
	ExitByday(c *BydayContext)

	// ExitWeekdaynum_options is called when exiting the weekdaynum_options production.
	ExitWeekdaynum_options(c *Weekdaynum_optionsContext)

	// ExitWeekdaynum is called when exiting the weekdaynum production.
	ExitWeekdaynum(c *WeekdaynumContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)

	// ExitByhour_clause is called when exiting the byhour_clause production.
	ExitByhour_clause(c *Byhour_clauseContext)

	// ExitHour_list is called when exiting the hour_list production.
	ExitHour_list(c *Hour_listContext)

	// ExitHour is called when exiting the hour production.
	ExitHour(c *HourContext)

	// ExitByminute_clause is called when exiting the byminute_clause production.
	ExitByminute_clause(c *Byminute_clauseContext)

	// ExitMinute_list is called when exiting the minute_list production.
	ExitMinute_list(c *Minute_listContext)

	// ExitMinute is called when exiting the minute production.
	ExitMinute(c *MinuteContext)

	// ExitBysecond_clause is called when exiting the bysecond_clause production.
	ExitBysecond_clause(c *Bysecond_clauseContext)

	// ExitSecond_list is called when exiting the second_list production.
	ExitSecond_list(c *Second_listContext)

	// ExitSecond is called when exiting the second production.
	ExitSecond(c *SecondContext)

	// ExitQuery_rewrite is called when exiting the query_rewrite production.
	ExitQuery_rewrite(c *Query_rewriteContext)

	// ExitBuild_clause is called when exiting the build_clause production.
	ExitBuild_clause(c *Build_clauseContext)

	// ExitMv_refresh_option is called when exiting the mv_refresh_option production.
	ExitMv_refresh_option(c *Mv_refresh_optionContext)

	// ExitMv_refresh_option_list is called when exiting the mv_refresh_option_list production.
	ExitMv_refresh_option_list(c *Mv_refresh_option_listContext)

	// ExitMv_refresh_clause is called when exiting the mv_refresh_clause production.
	ExitMv_refresh_clause(c *Mv_refresh_clauseContext)

	// ExitMv_log_purge_syn_asyn_clause is called when exiting the mv_log_purge_syn_asyn_clause production.
	ExitMv_log_purge_syn_asyn_clause(c *Mv_log_purge_syn_asyn_clauseContext)

	// ExitMv_log_purge_clause is called when exiting the mv_log_purge_clause production.
	ExitMv_log_purge_clause(c *Mv_log_purge_clauseContext)

	// ExitMv_log_with_syntax_item is called when exiting the mv_log_with_syntax_item production.
	ExitMv_log_with_syntax_item(c *Mv_log_with_syntax_itemContext)

	// ExitMv_log_with_syntax_item_list is called when exiting the mv_log_with_syntax_item_list production.
	ExitMv_log_with_syntax_item_list(c *Mv_log_with_syntax_item_listContext)

	// ExitMv_log_including_new_values is called when exiting the mv_log_including_new_values production.
	ExitMv_log_including_new_values(c *Mv_log_including_new_valuesContext)

	// ExitMv_log_with_clause_null is called when exiting the mv_log_with_clause_null production.
	ExitMv_log_with_clause_null(c *Mv_log_with_clause_nullContext)

	// ExitCreate_materialized_view_log_stmt is called when exiting the create_materialized_view_log_stmt production.
	ExitCreate_materialized_view_log_stmt(c *Create_materialized_view_log_stmtContext)

	// ExitPrebuilt_table_clause_null is called when exiting the prebuilt_table_clause_null production.
	ExitPrebuilt_table_clause_null(c *Prebuilt_table_clause_nullContext)

	// ExitCreate_materialized_view_stmt is called when exiting the create_materialized_view_stmt production.
	ExitCreate_materialized_view_stmt(c *Create_materialized_view_stmtContext)

	// ExitCreate_view_stmt is called when exiting the create_view_stmt production.
	ExitCreate_view_stmt(c *Create_view_stmtContext)

	// ExitCreate_view_stmt_body is called when exiting the create_view_stmt_body production.
	ExitCreate_view_stmt_body(c *Create_view_stmt_bodyContext)

	// ExitColumn_list3_options is called when exiting the column_list3_options production.
	ExitColumn_list3_options(c *Column_list3_optionsContext)

	// ExitColumn_list3 is called when exiting the column_list3 production.
	ExitColumn_list3(c *Column_list3Context)

	// ExitView_column_constraint_def is called when exiting the view_column_constraint_def production.
	ExitView_column_constraint_def(c *View_column_constraint_defContext)

	// ExitView_column_constraints is called when exiting the view_column_constraints production.
	ExitView_column_constraints(c *View_column_constraintsContext)

	// ExitView_column_constraint is called when exiting the view_column_constraint production.
	ExitView_column_constraint(c *View_column_constraintContext)

	// ExitView_column_constraint_action is called when exiting the view_column_constraint_action production.
	ExitView_column_constraint_action(c *View_column_constraint_actionContext)

	// ExitView_constraint_def is called when exiting the view_constraint_def production.
	ExitView_constraint_def(c *View_constraint_defContext)

	// ExitWith_schemabinding is called when exiting the with_schemabinding production.
	ExitWith_schemabinding(c *With_schemabindingContext)

	// ExitColumn_list_options is called when exiting the column_list_options production.
	ExitColumn_list_options(c *Column_list_optionsContext)

	// ExitWith_check_or_readonly_option is called when exiting the with_check_or_readonly_option production.
	ExitWith_check_or_readonly_option(c *With_check_or_readonly_optionContext)

	// ExitCheck_level_option is called when exiting the check_level_option production.
	ExitCheck_level_option(c *Check_level_optionContext)

	// ExitDecl_cursor is called when exiting the decl_cursor production.
	ExitDecl_cursor(c *Decl_cursorContext)

	// ExitDelete_stmt is called when exiting the delete_stmt production.
	ExitDelete_stmt(c *Delete_stmtContext)

	// ExitDelete_stmt_body is called when exiting the delete_stmt_body production.
	ExitDelete_stmt_body(c *Delete_stmt_bodyContext)

	// ExitDelete_multi_tv_option is called when exiting the delete_multi_tv_option production.
	ExitDelete_multi_tv_option(c *Delete_multi_tv_optionContext)

	// ExitCheck_limit_option is called when exiting the check_limit_option production.
	ExitCheck_limit_option(c *Check_limit_optionContext)

	// ExitWhere_current_option is called when exiting the where_current_option production.
	ExitWhere_current_option(c *Where_current_optionContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitStart_with_clause_null is called when exiting the start_with_clause_null production.
	ExitStart_with_clause_null(c *Start_with_clause_nullContext)

	// ExitConnect_by_item is called when exiting the connect_by_item production.
	ExitConnect_by_item(c *Connect_by_itemContext)

	// ExitConnect_by_clause is called when exiting the connect_by_clause production.
	ExitConnect_by_clause(c *Connect_by_clauseContext)

	// ExitHierarchical_query_clause is called when exiting the hierarchical_query_clause production.
	ExitHierarchical_query_clause(c *Hierarchical_query_clauseContext)

	// ExitNocycle_flag is called when exiting the nocycle_flag production.
	ExitNocycle_flag(c *Nocycle_flagContext)

	// ExitSearch_condition is called when exiting the search_condition production.
	ExitSearch_condition(c *Search_conditionContext)

	// ExitDisconnect_stmt is called when exiting the disconnect_stmt production.
	ExitDisconnect_stmt(c *Disconnect_stmtContext)

	// ExitDisconnect_option is called when exiting the disconnect_option production.
	ExitDisconnect_option(c *Disconnect_optionContext)

	// ExitDrop_stmt is called when exiting the drop_stmt production.
	ExitDrop_stmt(c *Drop_stmtContext)

	// ExitDrop_stmt_body_1 is called when exiting the drop_stmt_body_1 production.
	ExitDrop_stmt_body_1(c *Drop_stmt_body_1Context)

	// ExitDrop_stmt_2 is called when exiting the drop_stmt_2 production.
	ExitDrop_stmt_2(c *Drop_stmt_2Context)

	// ExitDrop_id_db_object is called when exiting the drop_id_db_object production.
	ExitDrop_id_db_object(c *Drop_id_db_objectContext)

	// ExitId_db_object is called when exiting the id_db_object production.
	ExitId_db_object(c *Id_db_objectContext)

	// ExitDrop_db_object is called when exiting the drop_db_object production.
	ExitDrop_db_object(c *Drop_db_objectContext)

	// ExitExist is called when exiting the exist production.
	ExitExist(c *ExistContext)

	// ExitNot_exist is called when exiting the not_exist production.
	ExitNot_exist(c *Not_existContext)

	// ExitDb_object is called when exiting the db_object production.
	ExitDb_object(c *Db_objectContext)

	// ExitIs_detach is called when exiting the is_detach production.
	ExitIs_detach(c *Is_detachContext)

	// ExitPurge_option is called when exiting the purge_option production.
	ExitPurge_option(c *Purge_optionContext)

	// ExitAlter_database_stmt is called when exiting the alter_database_stmt production.
	ExitAlter_database_stmt(c *Alter_database_stmtContext)

	// ExitAlter_system_action is called when exiting the alter_system_action production.
	ExitAlter_system_action(c *Alter_system_actionContext)

	// ExitAlter_database_action is called when exiting the alter_database_action production.
	ExitAlter_database_action(c *Alter_database_actionContext)

	// ExitForce is called when exiting the force production.
	ExitForce(c *ForceContext)

	// ExitTablespace_name is called when exiting the tablespace_name production.
	ExitTablespace_name(c *Tablespace_nameContext)

	// ExitRaft_name is called when exiting the raft_name production.
	ExitRaft_name(c *Raft_nameContext)

	// ExitFetch_into is called when exiting the fetch_into production.
	ExitFetch_into(c *Fetch_intoContext)

	// ExitBulk_or_single_into is called when exiting the bulk_or_single_into production.
	ExitBulk_or_single_into(c *Bulk_or_single_intoContext)

	// ExitFetch_stmt is called when exiting the fetch_stmt production.
	ExitFetch_stmt(c *Fetch_stmtContext)

	// ExitFetch_statement is called when exiting the fetch_statement production.
	ExitFetch_statement(c *Fetch_statementContext)

	// ExitFetch_tail is called when exiting the fetch_tail production.
	ExitFetch_tail(c *Fetch_tailContext)

	// ExitFetch_limit_option is called when exiting the fetch_limit_option production.
	ExitFetch_limit_option(c *Fetch_limit_optionContext)

	// ExitFetch_option is called when exiting the fetch_option production.
	ExitFetch_option(c *Fetch_optionContext)

	// ExitFetch_direct_option is called when exiting the fetch_direct_option production.
	ExitFetch_direct_option(c *Fetch_direct_optionContext)

	// ExitLog_errors_into is called when exiting the log_errors_into production.
	ExitLog_errors_into(c *Log_errors_intoContext)

	// ExitLog_errors_expression is called when exiting the log_errors_expression production.
	ExitLog_errors_expression(c *Log_errors_expressionContext)

	// ExitLog_errors_unlimited is called when exiting the log_errors_unlimited production.
	ExitLog_errors_unlimited(c *Log_errors_unlimitedContext)

	// ExitLog_errors is called when exiting the log_errors production.
	ExitLog_errors(c *Log_errorsContext)

	// ExitInsert_stmt is called when exiting the insert_stmt production.
	ExitInsert_stmt(c *Insert_stmtContext)

	// ExitInsert_stmt_body is called when exiting the insert_stmt_body production.
	ExitInsert_stmt_body(c *Insert_stmt_bodyContext)

	// ExitFull_column_list_options is called when exiting the full_column_list_options production.
	ExitFull_column_list_options(c *Full_column_list_optionsContext)

	// ExitIns_value_options is called when exiting the ins_value_options production.
	ExitIns_value_options(c *Ins_value_optionsContext)

	// ExitInsert_into_single is called when exiting the insert_into_single production.
	ExitInsert_into_single(c *Insert_into_singleContext)

	// ExitMulti_insert_into_list is called when exiting the multi_insert_into_list production.
	ExitMulti_insert_into_list(c *Multi_insert_into_listContext)

	// ExitMulti_insert_tag is called when exiting the multi_insert_tag production.
	ExitMulti_insert_tag(c *Multi_insert_tagContext)

	// ExitInsert_into_single_condition is called when exiting the insert_into_single_condition production.
	ExitInsert_into_single_condition(c *Insert_into_single_conditionContext)

	// ExitMulti_insert_into_condition_list is called when exiting the multi_insert_into_condition_list production.
	ExitMulti_insert_into_condition_list(c *Multi_insert_into_condition_listContext)

	// ExitMulti_insert_into_else is called when exiting the multi_insert_into_else production.
	ExitMulti_insert_into_else(c *Multi_insert_into_elseContext)

	// ExitMulti_insert_stmt_body is called when exiting the multi_insert_stmt_body production.
	ExitMulti_insert_stmt_body(c *Multi_insert_stmt_bodyContext)

	// ExitInsert_tail is called when exiting the insert_tail production.
	ExitInsert_tail(c *Insert_tailContext)

	// ExitInsert_action is called when exiting the insert_action production.
	ExitInsert_action(c *Insert_actionContext)

	// ExitRecord_var_values is called when exiting the record_var_values production.
	ExitRecord_var_values(c *Record_var_valuesContext)

	// ExitRecord_var_value is called when exiting the record_var_value production.
	ExitRecord_var_value(c *Record_var_valueContext)

	// ExitIns_value is called when exiting the ins_value production.
	ExitIns_value(c *Ins_valueContext)

	// ExitOpen_stmt is called when exiting the open_stmt production.
	ExitOpen_stmt(c *Open_stmtContext)

	// ExitOpen_statement is called when exiting the open_statement production.
	ExitOpen_statement(c *Open_statementContext)

	// ExitOpen_tail is called when exiting the open_tail production.
	ExitOpen_tail(c *Open_tailContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitRaise_stmt is called when exiting the raise_stmt production.
	ExitRaise_stmt(c *Raise_stmtContext)

	// ExitRollback_stmt is called when exiting the rollback_stmt production.
	ExitRollback_stmt(c *Rollback_stmtContext)

	// ExitTo_savepoint is called when exiting the to_savepoint production.
	ExitTo_savepoint(c *To_savepointContext)

	// ExitSavepoint_stmt is called when exiting the savepoint_stmt production.
	ExitSavepoint_stmt(c *Savepoint_stmtContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitAll_distinct_option is called when exiting the all_distinct_option production.
	ExitAll_distinct_option(c *All_distinct_optionContext)

	// ExitAll_distinct_option_2 is called when exiting the all_distinct_option_2 production.
	ExitAll_distinct_option_2(c *All_distinct_option_2Context)

	// ExitCorresponding_clause is called when exiting the corresponding_clause production.
	ExitCorresponding_clause(c *Corresponding_clauseContext)

	// ExitTop_option is called when exiting the top_option production.
	ExitTop_option(c *Top_optionContext)

	// ExitLimit_option is called when exiting the limit_option production.
	ExitLimit_option(c *Limit_optionContext)

	// ExitLimit_clause is called when exiting the limit_clause production.
	ExitLimit_clause(c *Limit_clauseContext)

	// ExitLimit_not_empty is called when exiting the limit_not_empty production.
	ExitLimit_not_empty(c *Limit_not_emptyContext)

	// ExitRow_limiting_clause is called when exiting the row_limiting_clause production.
	ExitRow_limiting_clause(c *Row_limiting_clauseContext)

	// ExitRow_num_desc is called when exiting the row_num_desc production.
	ExitRow_num_desc(c *Row_num_descContext)

	// ExitFirst_next_desc is called when exiting the first_next_desc production.
	ExitFirst_next_desc(c *First_next_descContext)

	// ExitSelect_item_list is called when exiting the select_item_list production.
	ExitSelect_item_list(c *Select_item_listContext)

	// ExitSelect_item is called when exiting the select_item production.
	ExitSelect_item(c *Select_itemContext)

	// ExitAs_alias is called when exiting the as_alias production.
	ExitAs_alias(c *As_aliasContext)

	// ExitSelect_tail is called when exiting the select_tail production.
	ExitSelect_tail(c *Select_tailContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitFrom_tv_list is called when exiting the from_tv_list production.
	ExitFrom_tv_list(c *From_tv_listContext)

	// ExitFrom_tv is called when exiting the from_tv production.
	ExitFrom_tv(c *From_tvContext)

	// ExitJoined_table is called when exiting the joined_table production.
	ExitJoined_table(c *Joined_tableContext)

	// ExitTrxid is called when exiting the trxid production.
	ExitTrxid(c *TrxidContext)

	// ExitFlashback_query_low is called when exiting the flashback_query_low production.
	ExitFlashback_query_low(c *Flashback_query_lowContext)

	// ExitTrxid_option is called when exiting the trxid_option production.
	ExitTrxid_option(c *Trxid_optionContext)

	// ExitRange_from_to is called when exiting the range_from_to production.
	ExitRange_from_to(c *Range_from_toContext)

	// ExitSample_exp is called when exiting the sample_exp production.
	ExitSample_exp(c *Sample_expContext)

	// ExitPivot_sfun is called when exiting the pivot_sfun production.
	ExitPivot_sfun(c *Pivot_sfunContext)

	// ExitPivot_sfun_lst is called when exiting the pivot_sfun_lst production.
	ExitPivot_sfun_lst(c *Pivot_sfun_lstContext)

	// ExitPivot_for_clause is called when exiting the pivot_for_clause production.
	ExitPivot_for_clause(c *Pivot_for_clauseContext)

	// ExitPivot_in_clause1_expr is called when exiting the pivot_in_clause1_expr production.
	ExitPivot_in_clause1_expr(c *Pivot_in_clause1_exprContext)

	// ExitPivot_in_clause_low_1 is called when exiting the pivot_in_clause_low_1 production.
	ExitPivot_in_clause_low_1(c *Pivot_in_clause_low_1Context)

	// ExitPivot_in_clause_low_2 is called when exiting the pivot_in_clause_low_2 production.
	ExitPivot_in_clause_low_2(c *Pivot_in_clause_low_2Context)

	// ExitPivot_in_clause_low is called when exiting the pivot_in_clause_low production.
	ExitPivot_in_clause_low(c *Pivot_in_clause_lowContext)

	// ExitPivot_xml is called when exiting the pivot_xml production.
	ExitPivot_xml(c *Pivot_xmlContext)

	// ExitPivot_clause_low is called when exiting the pivot_clause_low production.
	ExitPivot_clause_low(c *Pivot_clause_lowContext)

	// ExitUnpivot_val_col_lst is called when exiting the unpivot_val_col_lst production.
	ExitUnpivot_val_col_lst(c *Unpivot_val_col_lstContext)

	// ExitInclude_clause is called when exiting the include_clause production.
	ExitInclude_clause(c *Include_clauseContext)

	// ExitUnpivot_in_clause_expr is called when exiting the unpivot_in_clause_expr production.
	ExitUnpivot_in_clause_expr(c *Unpivot_in_clause_exprContext)

	// ExitUnpivot_in_clause_low is called when exiting the unpivot_in_clause_low production.
	ExitUnpivot_in_clause_low(c *Unpivot_in_clause_lowContext)

	// ExitUnpivot_clause_low is called when exiting the unpivot_clause_low production.
	ExitUnpivot_clause_low(c *Unpivot_clause_lowContext)

	// ExitSample_clause_low is called when exiting the sample_clause_low production.
	ExitSample_clause_low(c *Sample_clause_lowContext)

	// ExitNormal_tv_name is called when exiting the normal_tv_name production.
	ExitNormal_tv_name(c *Normal_tv_nameContext)

	// ExitNormal_tv_tail is called when exiting the normal_tv_tail production.
	ExitNormal_tv_tail(c *Normal_tv_tailContext)

	// ExitNormal_tv_tail_low is called when exiting the normal_tv_tail_low production.
	ExitNormal_tv_tail_low(c *Normal_tv_tail_lowContext)

	// ExitNormal_alias is called when exiting the normal_alias production.
	ExitNormal_alias(c *Normal_aliasContext)

	// ExitNormal_tv_tail_low2 is called when exiting the normal_tv_tail_low2 production.
	ExitNormal_tv_tail_low2(c *Normal_tv_tail_low2Context)

	// ExitNormal_tv_tail_low3 is called when exiting the normal_tv_tail_low3 production.
	ExitNormal_tv_tail_low3(c *Normal_tv_tail_low3Context)

	// ExitNormal_tv_derived_table_options is called when exiting the normal_tv_derived_table_options production.
	ExitNormal_tv_derived_table_options(c *Normal_tv_derived_table_optionsContext)

	// ExitNormal_tv_derived_table_low is called when exiting the normal_tv_derived_table_low production.
	ExitNormal_tv_derived_table_low(c *Normal_tv_derived_table_lowContext)

	// ExitNormal_tv_derived_table is called when exiting the normal_tv_derived_table production.
	ExitNormal_tv_derived_table(c *Normal_tv_derived_tableContext)

	// ExitSelect_with_paran_with_alias is called when exiting the select_with_paran_with_alias production.
	ExitSelect_with_paran_with_alias(c *Select_with_paran_with_aliasContext)

	// ExitFrom_table_exp is called when exiting the from_table_exp production.
	ExitFrom_table_exp(c *From_table_expContext)

	// ExitFrom_table_select_with_paran is called when exiting the from_table_select_with_paran production.
	ExitFrom_table_select_with_paran(c *From_table_select_with_paranContext)

	// ExitNormal_tv is called when exiting the normal_tv production.
	ExitNormal_tv(c *Normal_tvContext)

	// ExitXml_passing is called when exiting the xml_passing production.
	ExitXml_passing(c *Xml_passingContext)

	// ExitXmlcoldef_lst_options is called when exiting the xmlcoldef_lst_options production.
	ExitXmlcoldef_lst_options(c *Xmlcoldef_lst_optionsContext)

	// ExitXmlcoldef_lst is called when exiting the xmlcoldef_lst production.
	ExitXmlcoldef_lst(c *Xmlcoldef_lstContext)

	// ExitXmlcoldef is called when exiting the xmlcoldef production.
	ExitXmlcoldef(c *XmlcoldefContext)

	// ExitOn_error is called when exiting the on_error production.
	ExitOn_error(c *On_errorContext)

	// ExitJsoncol_lst is called when exiting the jsoncol_lst production.
	ExitJsoncol_lst(c *Jsoncol_lstContext)

	// ExitJsoncoldef_lst is called when exiting the jsoncoldef_lst production.
	ExitJsoncoldef_lst(c *Jsoncoldef_lstContext)

	// ExitJsoncoldef is called when exiting the jsoncoldef production.
	ExitJsoncoldef(c *JsoncoldefContext)

	// ExitJson_exists_col is called when exiting the json_exists_col production.
	ExitJson_exists_col(c *Json_exists_colContext)

	// ExitJson_qurey_col is called when exiting the json_qurey_col production.
	ExitJson_qurey_col(c *Json_qurey_colContext)

	// ExitJson_value_col is called when exiting the json_value_col production.
	ExitJson_value_col(c *Json_value_colContext)

	// ExitJson_nested_col is called when exiting the json_nested_col production.
	ExitJson_nested_col(c *Json_nested_colContext)

	// ExitOrdinality_col is called when exiting the ordinality_col production.
	ExitOrdinality_col(c *Ordinality_colContext)

	// ExitJoined_table_element is called when exiting the joined_table_element production.
	ExitJoined_table_element(c *Joined_table_elementContext)

	// ExitCross_outer_apply_clause is called when exiting the cross_outer_apply_clause production.
	ExitCross_outer_apply_clause(c *Cross_outer_apply_clauseContext)

	// ExitCross_outer_apply_join is called when exiting the cross_outer_apply_join production.
	ExitCross_outer_apply_join(c *Cross_outer_apply_joinContext)

	// ExitCross_join is called when exiting the cross_join production.
	ExitCross_join(c *Cross_joinContext)

	// ExitPartition_out_join is called when exiting the partition_out_join production.
	ExitPartition_out_join(c *Partition_out_joinContext)

	// ExitQualified_join is called when exiting the qualified_join production.
	ExitQualified_join(c *Qualified_joinContext)

	// ExitQualified_joinspec is called when exiting the qualified_joinspec production.
	ExitQualified_joinspec(c *Qualified_joinspecContext)

	// ExitNamed_columns_join is called when exiting the named_columns_join production.
	ExitNamed_columns_join(c *Named_columns_joinContext)

	// ExitJoin_hint is called when exiting the join_hint production.
	ExitJoin_hint(c *Join_hintContext)

	// ExitJoin_type is called when exiting the join_type production.
	ExitJoin_type(c *Join_typeContext)

	// ExitOuter_join_type is called when exiting the outer_join_type production.
	ExitOuter_join_type(c *Outer_join_typeContext)

	// ExitJoin_condition is called when exiting the join_condition production.
	ExitJoin_condition(c *Join_conditionContext)

	// ExitGroup_by_clause is called when exiting the group_by_clause production.
	ExitGroup_by_clause(c *Group_by_clauseContext)

	// ExitGroup_item is called when exiting the group_item production.
	ExitGroup_item(c *Group_itemContext)

	// ExitExp_rollup_cube_item2 is called when exiting the exp_rollup_cube_item2 production.
	ExitExp_rollup_cube_item2(c *Exp_rollup_cube_item2Context)

	// ExitExp_rollup_cube_item is called when exiting the exp_rollup_cube_item production.
	ExitExp_rollup_cube_item(c *Exp_rollup_cube_itemContext)

	// ExitGrouping_set_items is called when exiting the grouping_set_items production.
	ExitGrouping_set_items(c *Grouping_set_itemsContext)

	// ExitGrouping_set_item is called when exiting the grouping_set_item production.
	ExitGrouping_set_item(c *Grouping_set_itemContext)

	// ExitHaving_clause is called when exiting the having_clause production.
	ExitHaving_clause(c *Having_clauseContext)

	// ExitWithout_into_select is called when exiting the without_into_select production.
	ExitWithout_into_select(c *Without_into_selectContext)

	// ExitSel_clause_app is called when exiting the sel_clause_app production.
	ExitSel_clause_app(c *Sel_clause_appContext)

	// ExitSelect_clause is called when exiting the select_clause production.
	ExitSelect_clause(c *Select_clauseContext)

	// ExitSimple_select is called when exiting the simple_select production.
	ExitSimple_select(c *Simple_selectContext)

	// ExitSelect_with_paran is called when exiting the select_with_paran production.
	ExitSelect_with_paran(c *Select_with_paranContext)

	// ExitQuery_exp is called when exiting the query_exp production.
	ExitQuery_exp(c *Query_expContext)

	// ExitFor_xml_path is called when exiting the for_xml_path production.
	ExitFor_xml_path(c *For_xml_pathContext)

	// ExitWith_tag is called when exiting the with_tag production.
	ExitWith_tag(c *With_tagContext)

	// ExitWith_option is called when exiting the with_option production.
	ExitWith_option(c *With_optionContext)

	// ExitWith_clause is called when exiting the with_clause production.
	ExitWith_clause(c *With_clauseContext)

	// ExitWith_function_list is called when exiting the with_function_list production.
	ExitWith_function_list(c *With_function_listContext)

	// ExitFunc_def_inner is called when exiting the func_def_inner production.
	ExitFunc_def_inner(c *Func_def_innerContext)

	// ExitWith_function_list_element is called when exiting the with_function_list_element production.
	ExitWith_function_list_element(c *With_function_list_elementContext)

	// ExitWith_view_list is called when exiting the with_view_list production.
	ExitWith_view_list(c *With_view_listContext)

	// ExitDepth_type_option is called when exiting the depth_type_option production.
	ExitDepth_type_option(c *Depth_type_optionContext)

	// ExitSearch_clause is called when exiting the search_clause production.
	ExitSearch_clause(c *Search_clauseContext)

	// ExitCycle_clause is called when exiting the cycle_clause production.
	ExitCycle_clause(c *Cycle_clauseContext)

	// ExitWith_view_list_element is called when exiting the with_view_list_element production.
	ExitWith_view_list_element(c *With_view_list_elementContext)

	// ExitAssignment_obj_list is called when exiting the assignment_obj_list production.
	ExitAssignment_obj_list(c *Assignment_obj_listContext)

	// ExitAssignment_obj is called when exiting the assignment_obj production.
	ExitAssignment_obj(c *Assignment_objContext)

	// ExitOrder_by_options is called when exiting the order_by_options production.
	ExitOrder_by_options(c *Order_by_optionsContext)

	// ExitOrder_by is called when exiting the order_by production.
	ExitOrder_by(c *Order_byContext)

	// ExitAsc_desc_option is called when exiting the asc_desc_option production.
	ExitAsc_desc_option(c *Asc_desc_optionContext)

	// ExitNulls_last_option is called when exiting the nulls_last_option production.
	ExitNulls_last_option(c *Nulls_last_optionContext)

	// ExitCollate_option is called when exiting the collate_option production.
	ExitCollate_option(c *Collate_optionContext)

	// ExitOrder_by_list is called when exiting the order_by_list production.
	ExitOrder_by_list(c *Order_by_listContext)

	// ExitOrder_by_item is called when exiting the order_by_item production.
	ExitOrder_by_item(c *Order_by_itemContext)

	// ExitFor_update_options is called when exiting the for_update_options production.
	ExitFor_update_options(c *For_update_optionsContext)

	// ExitFor_update is called when exiting the for_update production.
	ExitFor_update(c *For_updateContext)

	// ExitSet_session_stmt is called when exiting the set_session_stmt production.
	ExitSet_session_stmt(c *Set_session_stmtContext)

	// ExitSet_trans_stmt is called when exiting the set_trans_stmt production.
	ExitSet_trans_stmt(c *Set_trans_stmtContext)

	// ExitTrans_mode_lstl is called when exiting the trans_mode_lstl production.
	ExitTrans_mode_lstl(c *Trans_mode_lstlContext)

	// ExitTrans_mode_lst is called when exiting the trans_mode_lst production.
	ExitTrans_mode_lst(c *Trans_mode_lstContext)

	// ExitTrans_mode is called when exiting the trans_mode production.
	ExitTrans_mode(c *Trans_modeContext)

	// ExitTime_zone_exp_new is called when exiting the time_zone_exp_new production.
	ExitTime_zone_exp_new(c *Time_zone_exp_newContext)

	// ExitTrans_rw_option is called when exiting the trans_rw_option production.
	ExitTrans_rw_option(c *Trans_rw_optionContext)

	// ExitTrans_level_option is called when exiting the trans_level_option production.
	ExitTrans_level_option(c *Trans_level_optionContext)

	// ExitLock_table_stmt is called when exiting the lock_table_stmt production.
	ExitLock_table_stmt(c *Lock_table_stmtContext)

	// ExitLock_mode_option is called when exiting the lock_mode_option production.
	ExitLock_mode_option(c *Lock_mode_optionContext)

	// ExitSet_identins_stmt is called when exiting the set_identins_stmt production.
	ExitSet_identins_stmt(c *Set_identins_stmtContext)

	// ExitSet_identins_option is called when exiting the set_identins_option production.
	ExitSet_identins_option(c *Set_identins_optionContext)

	// ExitTrunc_table_stmt is called when exiting the trunc_table_stmt production.
	ExitTrunc_table_stmt(c *Trunc_table_stmtContext)

	// ExitUpdate_stmt is called when exiting the update_stmt production.
	ExitUpdate_stmt(c *Update_stmtContext)

	// ExitUpdate_stmt_body is called when exiting the update_stmt_body production.
	ExitUpdate_stmt_body(c *Update_stmt_bodyContext)

	// ExitUpdate_tv_list is called when exiting the update_tv_list production.
	ExitUpdate_tv_list(c *Update_tv_listContext)

	// ExitReturn_item is called when exiting the return_item production.
	ExitReturn_item(c *Return_itemContext)

	// ExitReturn_item_list is called when exiting the return_item_list production.
	ExitReturn_item_list(c *Return_item_listContext)

	// ExitReturn_option is called when exiting the return_option production.
	ExitReturn_option(c *Return_optionContext)

	// ExitReturn_into_obj is called when exiting the return_into_obj production.
	ExitReturn_into_obj(c *Return_into_objContext)

	// ExitCollect_into_rset is called when exiting the collect_into_rset production.
	ExitCollect_into_rset(c *Collect_into_rsetContext)

	// ExitAlias_option is called when exiting the alias_option production.
	ExitAlias_option(c *Alias_optionContext)

	// ExitSet_value_list is called when exiting the set_value_list production.
	ExitSet_value_list(c *Set_value_listContext)

	// ExitSet_value_node is called when exiting the set_value_node production.
	ExitSet_value_node(c *Set_value_nodeContext)

	// ExitSet_col_list is called when exiting the set_col_list production.
	ExitSet_col_list(c *Set_col_listContext)

	// ExitUpdate_from_clause is called when exiting the update_from_clause production.
	ExitUpdate_from_clause(c *Update_from_clauseContext)

	// ExitMerge_into_stmt is called when exiting the merge_into_stmt production.
	ExitMerge_into_stmt(c *Merge_into_stmtContext)

	// ExitMerge_into_stmt_body is called when exiting the merge_into_stmt_body production.
	ExitMerge_into_stmt_body(c *Merge_into_stmt_bodyContext)

	// ExitMerge_into_sub_clause is called when exiting the merge_into_sub_clause production.
	ExitMerge_into_sub_clause(c *Merge_into_sub_clauseContext)

	// ExitMerge_update_clause is called when exiting the merge_update_clause production.
	ExitMerge_update_clause(c *Merge_update_clauseContext)

	// ExitMerge_insert_clause is called when exiting the merge_insert_clause production.
	ExitMerge_insert_clause(c *Merge_insert_clauseContext)

	// ExitCreate_profile_stmt is called when exiting the create_profile_stmt production.
	ExitCreate_profile_stmt(c *Create_profile_stmtContext)

	// ExitAlter_profile_stmt is called when exiting the alter_profile_stmt production.
	ExitAlter_profile_stmt(c *Alter_profile_stmtContext)

	// ExitCreate_user_stmt is called when exiting the create_user_stmt production.
	ExitCreate_user_stmt(c *Create_user_stmtContext)

	// ExitDefault_ts_name is called when exiting the default_ts_name production.
	ExitDefault_ts_name(c *Default_ts_nameContext)

	// ExitDefault_ts_name_lst is called when exiting the default_ts_name_lst production.
	ExitDefault_ts_name_lst(c *Default_ts_name_lstContext)

	// ExitDefault_ts_name_node is called when exiting the default_ts_name_node production.
	ExitDefault_ts_name_node(c *Default_ts_name_nodeContext)

	// ExitDefault_idx_ts_name is called when exiting the default_idx_ts_name production.
	ExitDefault_idx_ts_name(c *Default_idx_ts_nameContext)

	// ExitDefault_ts_name_low is called when exiting the default_ts_name_low production.
	ExitDefault_ts_name_low(c *Default_ts_name_lowContext)

	// ExitTemp_ts_name is called when exiting the temp_ts_name production.
	ExitTemp_ts_name(c *Temp_ts_nameContext)

	// ExitDefault_ts_group_name_low is called when exiting the default_ts_group_name_low production.
	ExitDefault_ts_group_name_low(c *Default_ts_group_name_lowContext)

	// ExitOn_schema is called when exiting the on_schema production.
	ExitOn_schema(c *On_schemaContext)

	// ExitReplace_old_pwd is called when exiting the replace_old_pwd production.
	ExitReplace_old_pwd(c *Replace_old_pwdContext)

	// ExitAlter_user_stmt is called when exiting the alter_user_stmt production.
	ExitAlter_user_stmt(c *Alter_user_stmtContext)

	// ExitUser_encrypt_options is called when exiting the user_encrypt_options production.
	ExitUser_encrypt_options(c *User_encrypt_optionsContext)

	// ExitUser_encrypt_option is called when exiting the user_encrypt_option production.
	ExitUser_encrypt_option(c *User_encrypt_optionContext)

	// ExitAuthent_type_options is called when exiting the authent_type_options production.
	ExitAuthent_type_options(c *Authent_type_optionsContext)

	// ExitHash_cipher_option is called when exiting the hash_cipher_option production.
	ExitHash_cipher_option(c *Hash_cipher_optionContext)

	// ExitAuthent_type is called when exiting the authent_type production.
	ExitAuthent_type(c *Authent_typeContext)

	// ExitForce_format is called when exiting the force_format production.
	ExitForce_format(c *Force_formatContext)

	// ExitAs is called when exiting the as production.
	ExitAs(c *AsContext)

	// ExitPwd_policy is called when exiting the pwd_policy production.
	ExitPwd_policy(c *Pwd_policyContext)

	// ExitAccount_lock is called when exiting the account_lock production.
	ExitAccount_lock(c *Account_lockContext)

	// ExitRead_only_flag is called when exiting the read_only_flag production.
	ExitRead_only_flag(c *Read_only_flagContext)

	// ExitRead_only_flag_not_null is called when exiting the read_only_flag_not_null production.
	ExitRead_only_flag_not_null(c *Read_only_flag_not_nullContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitExpire is called when exiting the expire production.
	ExitExpire(c *ExpireContext)

	// ExitResource_limit_options is called when exiting the resource_limit_options production.
	ExitResource_limit_options(c *Resource_limit_optionsContext)

	// ExitResource_limit_list is called when exiting the resource_limit_list production.
	ExitResource_limit_list(c *Resource_limit_listContext)

	// ExitResource_limit_list_with_comma is called when exiting the resource_limit_list_with_comma production.
	ExitResource_limit_list_with_comma(c *Resource_limit_list_with_commaContext)

	// ExitResource_limit_list_with_empty is called when exiting the resource_limit_list_with_empty production.
	ExitResource_limit_list_with_empty(c *Resource_limit_list_with_emptyContext)

	// ExitResource_limit is called when exiting the resource_limit production.
	ExitResource_limit(c *Resource_limitContext)

	// ExitResource_limit_value is called when exiting the resource_limit_value production.
	ExitResource_limit_value(c *Resource_limit_valueContext)

	// ExitCreate_audit_rule_stmt is called when exiting the create_audit_rule_stmt production.
	ExitCreate_audit_rule_stmt(c *Create_audit_rule_stmtContext)

	// ExitRule_name is called when exiting the rule_name production.
	ExitRule_name(c *Rule_nameContext)

	// ExitAudit_rule_action is called when exiting the audit_rule_action production.
	ExitAudit_rule_action(c *Audit_rule_actionContext)

	// ExitBy_login_or_all is called when exiting the by_login_or_all production.
	ExitBy_login_or_all(c *By_login_or_allContext)

	// ExitAllow_ip_list is called when exiting the allow_ip_list production.
	ExitAllow_ip_list(c *Allow_ip_listContext)

	// ExitNot_allow_ip_list is called when exiting the not_allow_ip_list production.
	ExitNot_allow_ip_list(c *Not_allow_ip_listContext)

	// ExitIp_list is called when exiting the ip_list production.
	ExitIp_list(c *Ip_listContext)

	// ExitAllow_dt_list is called when exiting the allow_dt_list production.
	ExitAllow_dt_list(c *Allow_dt_listContext)

	// ExitNot_allow_dt_list is called when exiting the not_allow_dt_list production.
	ExitNot_allow_dt_list(c *Not_allow_dt_listContext)

	// ExitDt_list is called when exiting the dt_list production.
	ExitDt_list(c *Dt_listContext)

	// ExitDt is called when exiting the dt production.
	ExitDt(c *DtContext)

	// ExitOp_freq is called when exiting the op_freq production.
	ExitOp_freq(c *Op_freqContext)

	// ExitQuota_val is called when exiting the quota_val production.
	ExitQuota_val(c *Quota_valContext)

	// ExitQuota_ts_node is called when exiting the quota_ts_node production.
	ExitQuota_ts_node(c *Quota_ts_nodeContext)

	// ExitQuota_ts_lst is called when exiting the quota_ts_lst production.
	ExitQuota_ts_lst(c *Quota_ts_lstContext)

	// ExitQuota_ts is called when exiting the quota_ts production.
	ExitQuota_ts(c *Quota_tsContext)

	// ExitCreate_role_stmt is called when exiting the create_role_stmt production.
	ExitCreate_role_stmt(c *Create_role_stmtContext)

	// ExitCreate_dblink_stmt is called when exiting the create_dblink_stmt production.
	ExitCreate_dblink_stmt(c *Create_dblink_stmtContext)

	// ExitDb_type_str is called when exiting the db_type_str production.
	ExitDb_type_str(c *Db_type_strContext)

	// ExitDblink_option_lst_options is called when exiting the dblink_option_lst_options production.
	ExitDblink_option_lst_options(c *Dblink_option_lst_optionsContext)

	// ExitDblink_option_lst is called when exiting the dblink_option_lst production.
	ExitDblink_option_lst(c *Dblink_option_lstContext)

	// ExitDblink_option is called when exiting the dblink_option production.
	ExitDblink_option(c *Dblink_optionContext)

	// ExitCreate_synonym_stmt is called when exiting the create_synonym_stmt production.
	ExitCreate_synonym_stmt(c *Create_synonym_stmtContext)

	// ExitFull_synonym_name is called when exiting the full_synonym_name production.
	ExitFull_synonym_name(c *Full_synonym_nameContext)

	// ExitFull_obj_name is called when exiting the full_obj_name production.
	ExitFull_obj_name(c *Full_obj_nameContext)

	// ExitCreate_domain_stmt is called when exiting the create_domain_stmt production.
	ExitCreate_domain_stmt(c *Create_domain_stmtContext)

	// ExitDomain_default_option is called when exiting the domain_default_option production.
	ExitDomain_default_option(c *Domain_default_optionContext)

	// ExitDomain_constraints_option is called when exiting the domain_constraints_option production.
	ExitDomain_constraints_option(c *Domain_constraints_optionContext)

	// ExitDomain_constraints_def is called when exiting the domain_constraints_def production.
	ExitDomain_constraints_def(c *Domain_constraints_defContext)

	// ExitDomain_constraints is called when exiting the domain_constraints production.
	ExitDomain_constraints(c *Domain_constraintsContext)

	// ExitDomain_constraint is called when exiting the domain_constraint production.
	ExitDomain_constraint(c *Domain_constraintContext)

	// ExitDomain_constraint_name_def_options is called when exiting the domain_constraint_name_def_options production.
	ExitDomain_constraint_name_def_options(c *Domain_constraint_name_def_optionsContext)

	// ExitDomain_constraint_name_def is called when exiting the domain_constraint_name_def production.
	ExitDomain_constraint_name_def(c *Domain_constraint_name_defContext)

	// ExitDomain_constraint_name is called when exiting the domain_constraint_name production.
	ExitDomain_constraint_name(c *Domain_constraint_nameContext)

	// ExitCreate_character_set_stmt is called when exiting the create_character_set_stmt production.
	ExitCreate_character_set_stmt(c *Create_character_set_stmtContext)

	// ExitCharacter_set_source is called when exiting the character_set_source production.
	ExitCharacter_set_source(c *Character_set_sourceContext)

	// ExitExisting_character_set_name is called when exiting the existing_character_set_name production.
	ExitExisting_character_set_name(c *Existing_character_set_nameContext)

	// ExitCharacter_set_name is called when exiting the character_set_name production.
	ExitCharacter_set_name(c *Character_set_nameContext)

	// ExitCollate_clause_option is called when exiting the collate_clause_option production.
	ExitCollate_clause_option(c *Collate_clause_optionContext)

	// ExitCollation_name is called when exiting the collation_name production.
	ExitCollation_name(c *Collation_nameContext)

	// ExitCreate_collation_stmt is called when exiting the create_collation_stmt production.
	ExitCreate_collation_stmt(c *Create_collation_stmtContext)

	// ExitExisting_collation_name is called when exiting the existing_collation_name production.
	ExitExisting_collation_name(c *Existing_collation_nameContext)

	// ExitPad_option is called when exiting the pad_option production.
	ExitPad_option(c *Pad_optionContext)

	// ExitCreate_sequence_stmt is called when exiting the create_sequence_stmt production.
	ExitCreate_sequence_stmt(c *Create_sequence_stmtContext)

	// ExitSequence_option_list_options is called when exiting the sequence_option_list_options production.
	ExitSequence_option_list_options(c *Sequence_option_list_optionsContext)

	// ExitSequence_option_list is called when exiting the sequence_option_list production.
	ExitSequence_option_list(c *Sequence_option_listContext)

	// ExitSequence_option is called when exiting the sequence_option production.
	ExitSequence_option(c *Sequence_optionContext)

	// ExitSequence_name is called when exiting the sequence_name production.
	ExitSequence_name(c *Sequence_nameContext)

	// ExitIncrement_option is called when exiting the increment_option production.
	ExitIncrement_option(c *Increment_optionContext)

	// ExitStart_option is called when exiting the start_option production.
	ExitStart_option(c *Start_optionContext)

	// ExitCurrent_option is called when exiting the current_option production.
	ExitCurrent_option(c *Current_optionContext)

	// ExitMaxvalue_option is called when exiting the maxvalue_option production.
	ExitMaxvalue_option(c *Maxvalue_optionContext)

	// ExitMinvalue_option is called when exiting the minvalue_option production.
	ExitMinvalue_option(c *Minvalue_optionContext)

	// ExitCycle_option is called when exiting the cycle_option production.
	ExitCycle_option(c *Cycle_optionContext)

	// ExitCache_option is called when exiting the cache_option production.
	ExitCache_option(c *Cache_optionContext)

	// ExitOrder_option is called when exiting the order_option production.
	ExitOrder_option(c *Order_optionContext)

	// ExitSeq_local_option is called when exiting the seq_local_option production.
	ExitSeq_local_option(c *Seq_local_optionContext)

	// ExitWhenever_stmt_options is called when exiting the whenever_stmt_options production.
	ExitWhenever_stmt_options(c *Whenever_stmt_optionsContext)

	// ExitWhenever_stmt is called when exiting the whenever_stmt production.
	ExitWhenever_stmt(c *Whenever_stmtContext)

	// ExitGrant_stmt is called when exiting the grant_stmt production.
	ExitGrant_stmt(c *Grant_stmtContext)

	// ExitGrant_tag is called when exiting the grant_tag production.
	ExitGrant_tag(c *Grant_tagContext)

	// ExitGrant_stmt_body is called when exiting the grant_stmt_body production.
	ExitGrant_stmt_body(c *Grant_stmt_bodyContext)

	// ExitRevoke_stmt is called when exiting the revoke_stmt production.
	ExitRevoke_stmt(c *Revoke_stmtContext)

	// ExitRevoke_stmt_body is called when exiting the revoke_stmt_body production.
	ExitRevoke_stmt_body(c *Revoke_stmt_bodyContext)

	// ExitGrant_privs is called when exiting the grant_privs production.
	ExitGrant_privs(c *Grant_privsContext)

	// ExitGrant_priv_list is called when exiting the grant_priv_list production.
	ExitGrant_priv_list(c *Grant_priv_listContext)

	// ExitGrant_priv_off is called when exiting the grant_priv_off production.
	ExitGrant_priv_off(c *Grant_priv_offContext)

	// ExitGrant_priv is called when exiting the grant_priv production.
	ExitGrant_priv(c *Grant_privContext)

	// ExitRevoke_action is called when exiting the revoke_action production.
	ExitRevoke_action(c *Revoke_actionContext)

	// ExitDb_priv_list is called when exiting the db_priv_list production.
	ExitDb_priv_list(c *Db_priv_listContext)

	// ExitDb_priv is called when exiting the db_priv production.
	ExitDb_priv(c *Db_privContext)

	// ExitBy_grantor is called when exiting the by_grantor production.
	ExitBy_grantor(c *By_grantorContext)

	// ExitGrantees is called when exiting the grantees production.
	ExitGrantees(c *GranteesContext)

	// ExitCreate_schema_stmt is called when exiting the create_schema_stmt production.
	ExitCreate_schema_stmt(c *Create_schema_stmtContext)

	// ExitOprt_arg is called when exiting the oprt_arg production.
	ExitOprt_arg(c *Oprt_argContext)

	// ExitOprt_arg_lst is called when exiting the oprt_arg_lst production.
	ExitOprt_arg_lst(c *Oprt_arg_lstContext)

	// ExitCreate_operator_stmt is called when exiting the create_operator_stmt production.
	ExitCreate_operator_stmt(c *Create_operator_stmtContext)

	// ExitDrop_operator_stmt is called when exiting the drop_operator_stmt production.
	ExitDrop_operator_stmt(c *Drop_operator_stmtContext)

	// ExitGrant_and_ddl is called when exiting the grant_and_ddl production.
	ExitGrant_and_ddl(c *Grant_and_ddlContext)

	// ExitTop_exp is called when exiting the top_exp production.
	ExitTop_exp(c *Top_expContext)

	// ExitU_oprt is called when exiting the u_oprt production.
	ExitU_oprt(c *U_oprtContext)

	// ExitQualified_u_oprt is called when exiting the qualified_u_oprt production.
	ExitQualified_u_oprt(c *Qualified_u_oprtContext)

	// ExitExp_u_oprt is called when exiting the exp_u_oprt production.
	ExitExp_u_oprt(c *Exp_u_oprtContext)

	// ExitRaw_exp is called when exiting the raw_exp production.
	ExitRaw_exp(c *Raw_expContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitFrom_first_last_option is called when exiting the from_first_last_option production.
	ExitFrom_first_last_option(c *From_first_last_optionContext)

	// ExitAfun_arg_lst is called when exiting the afun_arg_lst production.
	ExitAfun_arg_lst(c *Afun_arg_lstContext)

	// ExitAfun_arg_lst_low is called when exiting the afun_arg_lst_low production.
	ExitAfun_arg_lst_low(c *Afun_arg_lst_lowContext)

	// ExitIn_value_exp is called when exiting the in_value_exp production.
	ExitIn_value_exp(c *In_value_expContext)

	// ExitAfun_partition_by is called when exiting the afun_partition_by production.
	ExitAfun_partition_by(c *Afun_partition_byContext)

	// ExitAfun_windowing is called when exiting the afun_windowing production.
	ExitAfun_windowing(c *Afun_windowingContext)

	// ExitAfun_windowing_type is called when exiting the afun_windowing_type production.
	ExitAfun_windowing_type(c *Afun_windowing_typeContext)

	// ExitAfun_range_clause is called when exiting the afun_range_clause production.
	ExitAfun_range_clause(c *Afun_range_clauseContext)

	// ExitPexp is called when exiting the pexp production.
	ExitPexp(c *PexpContext)

	// ExitPexp_pfx is called when exiting the pexp_pfx production.
	ExitPexp_pfx(c *Pexp_pfxContext)

	// ExitPexp_cast is called when exiting the pexp_cast production.
	ExitPexp_cast(c *Pexp_castContext)

	// ExitPexp_b is called when exiting the pexp_b production.
	ExitPexp_b(c *Pexp_bContext)

	// ExitPexp_a is called when exiting the pexp_a production.
	ExitPexp_a(c *Pexp_aContext)

	// ExitPexp_c is called when exiting the pexp_c production.
	ExitPexp_c(c *Pexp_cContext)

	// ExitPexp_c_insert is called when exiting the pexp_c_insert production.
	ExitPexp_c_insert(c *Pexp_c_insertContext)

	// ExitPexp_d is called when exiting the pexp_d production.
	ExitPexp_d(c *Pexp_dContext)

	// ExitPexp_e is called when exiting the pexp_e production.
	ExitPexp_e(c *Pexp_eContext)

	// ExitPexp_pfx2 is called when exiting the pexp_pfx2 production.
	ExitPexp_pfx2(c *Pexp_pfx2Context)

	// ExitMember2 is called when exiting the member2 production.
	ExitMember2(c *Member2Context)

	// ExitPexp_c2_insert is called when exiting the pexp_c2_insert production.
	ExitPexp_c2_insert(c *Pexp_c2_insertContext)

	// ExitMember_access2 is called when exiting the member_access2 production.
	ExitMember_access2(c *Member_access2Context)

	// ExitInvocation_expression2 is called when exiting the invocation_expression2 production.
	ExitInvocation_expression2(c *Invocation_expression2Context)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitMember_access is called when exiting the member_access production.
	ExitMember_access(c *Member_accessContext)

	// ExitInvocation_expression is called when exiting the invocation_expression production.
	ExitInvocation_expression(c *Invocation_expressionContext)

	// ExitInvocation_expression_low is called when exiting the invocation_expression_low production.
	ExitInvocation_expression_low(c *Invocation_expression_lowContext)

	// ExitXmlagg_inv_expression is called when exiting the xmlagg_inv_expression production.
	ExitXmlagg_inv_expression(c *Xmlagg_inv_expressionContext)

	// ExitXmlfun_inv_expression is called when exiting the xmlfun_inv_expression production.
	ExitXmlfun_inv_expression(c *Xmlfun_inv_expressionContext)

	// ExitXmlele_name is called when exiting the xmlele_name production.
	ExitXmlele_name(c *Xmlele_nameContext)

	// ExitXmlele_sub_lst is called when exiting the xmlele_sub_lst production.
	ExitXmlele_sub_lst(c *Xmlele_sub_lstContext)

	// ExitXmlattr_lst is called when exiting the xmlattr_lst production.
	ExitXmlattr_lst(c *Xmlattr_lstContext)

	// ExitXmlattr is called when exiting the xmlattr production.
	ExitXmlattr(c *XmlattrContext)

	// ExitXmlval_lst is called when exiting the xmlval_lst production.
	ExitXmlval_lst(c *Xmlval_lstContext)

	// ExitKeep_clause is called when exiting the keep_clause production.
	ExitKeep_clause(c *Keep_clauseContext)

	// ExitWithin_clause is called when exiting the within_clause production.
	ExitWithin_clause(c *Within_clauseContext)

	// ExitTypeof_expression is called when exiting the typeof_expression production.
	ExitTypeof_expression(c *Typeof_expressionContext)

	// ExitNew_obj_expression is called when exiting the new_obj_expression production.
	ExitNew_obj_expression(c *New_obj_expressionContext)

	// ExitNew_arr_expression is called when exiting the new_arr_expression production.
	ExitNew_arr_expression(c *New_arr_expressionContext)

	// ExitArray_creation_expression is called when exiting the array_creation_expression production.
	ExitArray_creation_expression(c *Array_creation_expressionContext)

	// ExitPlsql_datatype_ex is called when exiting the plsql_datatype_ex production.
	ExitPlsql_datatype_ex(c *Plsql_datatype_exContext)

	// ExitNew_array_type is called when exiting the new_array_type production.
	ExitNew_array_type(c *New_array_typeContext)

	// ExitOpt_array_initializer is called when exiting the opt_array_initializer production.
	ExitOpt_array_initializer(c *Opt_array_initializerContext)

	// ExitArray_initializer is called when exiting the array_initializer production.
	ExitArray_initializer(c *Array_initializerContext)

	// ExitVariable_initializer_list is called when exiting the variable_initializer_list production.
	ExitVariable_initializer_list(c *Variable_initializer_listContext)

	// ExitVariable_initializer is called when exiting the variable_initializer production.
	ExitVariable_initializer(c *Variable_initializerContext)

	// ExitOpt_comma is called when exiting the opt_comma production.
	ExitOpt_comma(c *Opt_commaContext)

	// ExitSizeof_expression is called when exiting the sizeof_expression production.
	ExitSizeof_expression(c *Sizeof_expressionContext)

	// ExitElement_access is called when exiting the element_access production.
	ExitElement_access(c *Element_accessContext)

	// ExitDecode_case is called when exiting the decode_case production.
	ExitDecode_case(c *Decode_caseContext)

	// ExitElse_exp is called when exiting the else_exp production.
	ExitElse_exp(c *Else_expContext)

	// ExitBoolean_case is called when exiting the boolean_case production.
	ExitBoolean_case(c *Boolean_caseContext)

	// ExitIf_exp is called when exiting the if_exp production.
	ExitIf_exp(c *If_expContext)

	// ExitBool_when_list is called when exiting the bool_when_list production.
	ExitBool_when_list(c *Bool_when_listContext)

	// ExitOps is called when exiting the ops production.
	ExitOps(c *OpsContext)

	// ExitValue_list is called when exiting the value_list production.
	ExitValue_list(c *Value_listContext)

	// ExitIn_value_list is called when exiting the in_value_list production.
	ExitIn_value_list(c *In_value_listContext)

	// ExitValue_list_set is called when exiting the value_list_set production.
	ExitValue_list_set(c *Value_list_setContext)

	// ExitComma_list is called when exiting the comma_list production.
	ExitComma_list(c *Comma_listContext)

	// ExitIns_value_list is called when exiting the ins_value_list production.
	ExitIns_value_list(c *Ins_value_listContext)

	// ExitNull_value is called when exiting the null_value production.
	ExitNull_value(c *Null_valueContext)

	// ExitId_and_rsvd_word_others is called when exiting the id_and_rsvd_word_others production.
	ExitId_and_rsvd_word_others(c *Id_and_rsvd_word_othersContext)

	// ExitId_and_rsvd_word is called when exiting the id_and_rsvd_word production.
	ExitId_and_rsvd_word(c *Id_and_rsvd_wordContext)

	// ExitStm_param is called when exiting the stm_param production.
	ExitStm_param(c *Stm_paramContext)

	// ExitStm_param_normal is called when exiting the stm_param_normal production.
	ExitStm_param_normal(c *Stm_param_normalContext)

	// ExitStm_param_name is called when exiting the stm_param_name production.
	ExitStm_param_name(c *Stm_param_nameContext)

	// ExitParam_name_options is called when exiting the param_name_options production.
	ExitParam_name_options(c *Param_name_optionsContext)

	// ExitContains_query_exp is called when exiting the contains_query_exp production.
	ExitContains_query_exp(c *Contains_query_expContext)

	// ExitContains_query_exp_lst is called when exiting the contains_query_exp_lst production.
	ExitContains_query_exp_lst(c *Contains_query_exp_lstContext)

	// ExitContains_exp is called when exiting the contains_exp production.
	ExitContains_exp(c *Contains_expContext)

	// ExitStrict_option is called when exiting the strict_option production.
	ExitStrict_option(c *Strict_optionContext)

	// ExitWith_unique_option is called when exiting the with_unique_option production.
	ExitWith_unique_option(c *With_unique_optionContext)

	// ExitType_option is called when exiting the type_option production.
	ExitType_option(c *Type_optionContext)

	// ExitType_element is called when exiting the type_element production.
	ExitType_element(c *Type_elementContext)

	// ExitType_element_list is called when exiting the type_element_list production.
	ExitType_element_list(c *Type_element_listContext)

	// ExitBool_exp is called when exiting the bool_exp production.
	ExitBool_exp(c *Bool_expContext)

	// ExitBool_exp_element is called when exiting the bool_exp_element production.
	ExitBool_exp_element(c *Bool_exp_elementContext)

	// ExitQuery_any_options is called when exiting the query_any_options production.
	ExitQuery_any_options(c *Query_any_optionsContext)

	// ExitGlobal_var is called when exiting the global_var production.
	ExitGlobal_var(c *Global_varContext)

	// ExitReserved_word is called when exiting the reserved_word production.
	ExitReserved_word(c *Reserved_wordContext)

	// ExitNew_none_reserved_word is called when exiting the new_none_reserved_word production.
	ExitNew_none_reserved_word(c *New_none_reserved_wordContext)

	// ExitInterval_nresvd_word is called when exiting the interval_nresvd_word production.
	ExitInterval_nresvd_word(c *Interval_nresvd_wordContext)

	// ExitVariable_resvd_word is called when exiting the variable_resvd_word production.
	ExitVariable_resvd_word(c *Variable_resvd_wordContext)

	// ExitAlias_resvd_word is called when exiting the alias_resvd_word production.
	ExitAlias_resvd_word(c *Alias_resvd_wordContext)

	// ExitSchname_resvd_word is called when exiting the schname_resvd_word production.
	ExitSchname_resvd_word(c *Schname_resvd_wordContext)

	// ExitRaw_id is called when exiting the raw_id production.
	ExitRaw_id(c *Raw_idContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitQualified_name is called when exiting the qualified_name production.
	ExitQualified_name(c *Qualified_nameContext)

	// ExitQualified_name2 is called when exiting the qualified_name2 production.
	ExitQualified_name2(c *Qualified_name2Context)

	// ExitVariable_name is called when exiting the variable_name production.
	ExitVariable_name(c *Variable_nameContext)

	// ExitEnd_loop_label_null is called when exiting the end_loop_label_null production.
	ExitEnd_loop_label_null(c *End_loop_label_nullContext)

	// ExitLabel_name_options is called when exiting the label_name_options production.
	ExitLabel_name_options(c *Label_name_optionsContext)

	// ExitLabel_name is called when exiting the label_name production.
	ExitLabel_name(c *Label_nameContext)

	// ExitDatabase_name is called when exiting the database_name production.
	ExitDatabase_name(c *Database_nameContext)

	// ExitBackup_name is called when exiting the backup_name production.
	ExitBackup_name(c *Backup_nameContext)

	// ExitFull_proc_name is called when exiting the full_proc_name production.
	ExitFull_proc_name(c *Full_proc_nameContext)

	// ExitFull_proc_name2 is called when exiting the full_proc_name2 production.
	ExitFull_proc_name2(c *Full_proc_name2Context)

	// ExitFull_fun_name is called when exiting the full_fun_name production.
	ExitFull_fun_name(c *Full_fun_nameContext)

	// ExitFull_table_name is called when exiting the full_table_name production.
	ExitFull_table_name(c *Full_table_nameContext)

	// ExitFull_grp_name is called when exiting the full_grp_name production.
	ExitFull_grp_name(c *Full_grp_nameContext)

	// ExitFull_table_name2 is called when exiting the full_table_name2 production.
	ExitFull_table_name2(c *Full_table_name2Context)

	// ExitFull_partition_name is called when exiting the full_partition_name production.
	ExitFull_partition_name(c *Full_partition_nameContext)

	// ExitFull_schema_name is called when exiting the full_schema_name production.
	ExitFull_schema_name(c *Full_schema_nameContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitConstraint_name is called when exiting the constraint_name production.
	ExitConstraint_name(c *Constraint_nameContext)

	// ExitFull_trigger_name is called when exiting the full_trigger_name production.
	ExitFull_trigger_name(c *Full_trigger_nameContext)

	// ExitFull_trigger_name2 is called when exiting the full_trigger_name2 production.
	ExitFull_trigger_name2(c *Full_trigger_name2Context)

	// ExitFull_view_name is called when exiting the full_view_name production.
	ExitFull_view_name(c *Full_view_nameContext)

	// ExitFull_view_name2 is called when exiting the full_view_name2 production.
	ExitFull_view_name2(c *Full_view_name2Context)

	// ExitCursor_name is called when exiting the cursor_name production.
	ExitCursor_name(c *Cursor_nameContext)

	// ExitTrigger_name is called when exiting the trigger_name production.
	ExitTrigger_name(c *Trigger_nameContext)

	// ExitLogin_name is called when exiting the login_name production.
	ExitLogin_name(c *Login_nameContext)

	// ExitProfile_name is called when exiting the profile_name production.
	ExitProfile_name(c *Profile_nameContext)

	// ExitUser_name is called when exiting the user_name production.
	ExitUser_name(c *User_nameContext)

	// ExitRole_name is called when exiting the role_name production.
	ExitRole_name(c *Role_nameContext)

	// ExitUser_role_name is called when exiting the user_role_name production.
	ExitUser_role_name(c *User_role_nameContext)

	// ExitRole_name_list is called when exiting the role_name_list production.
	ExitRole_name_list(c *Role_name_listContext)

	// ExitFull_func_name is called when exiting the full_func_name production.
	ExitFull_func_name(c *Full_func_nameContext)

	// ExitParam_name is called when exiting the param_name production.
	ExitParam_name(c *Param_nameContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitIndex_name2 is called when exiting the index_name2 production.
	ExitIndex_name2(c *Index_name2Context)

	// ExitTrig_old_name is called when exiting the trig_old_name production.
	ExitTrig_old_name(c *Trig_old_nameContext)

	// ExitTrig_new_name is called when exiting the trig_new_name production.
	ExitTrig_new_name(c *Trig_new_nameContext)

	// ExitFull_tv_name is called when exiting the full_tv_name production.
	ExitFull_tv_name(c *Full_tv_nameContext)

	// ExitFull_object_name is called when exiting the full_object_name production.
	ExitFull_object_name(c *Full_object_nameContext)

	// ExitOrient_option is called when exiting the orient_option production.
	ExitOrient_option(c *Orient_optionContext)

	// ExitDatepart is called when exiting the datepart production.
	ExitDatepart(c *DatepartContext)

	// ExitDatepart_op is called when exiting the datepart_op production.
	ExitDatepart_op(c *Datepart_opContext)

	// ExitDatead_fun is called when exiting the datead_fun production.
	ExitDatead_fun(c *Datead_funContext)

	// ExitReturning is called when exiting the returning production.
	ExitReturning(c *ReturningContext)

	// ExitPretty is called when exiting the pretty production.
	ExitPretty(c *PrettyContext)

	// ExitWrapper_flag is called when exiting the wrapper_flag production.
	ExitWrapper_flag(c *Wrapper_flagContext)

	// ExitArray_wrapper is called when exiting the array_wrapper production.
	ExitArray_wrapper(c *Array_wrapperContext)

	// ExitJson_tail_on_empty is called when exiting the json_tail_on_empty production.
	ExitJson_tail_on_empty(c *Json_tail_on_emptyContext)

	// ExitEmpty_handle is called when exiting the empty_handle production.
	ExitEmpty_handle(c *Empty_handleContext)

	// ExitJson_tail_on_error_null is called when exiting the json_tail_on_error_null production.
	ExitJson_tail_on_error_null(c *Json_tail_on_error_nullContext)

	// ExitError_handle is called when exiting the error_handle production.
	ExitError_handle(c *Error_handleContext)

	// ExitSavepoint_name is called when exiting the savepoint_name production.
	ExitSavepoint_name(c *Savepoint_nameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitAlias_2 is called when exiting the alias_2 production.
	ExitAlias_2(c *Alias_2Context)

	// ExitFull_column_name is called when exiting the full_column_name production.
	ExitFull_column_name(c *Full_column_nameContext)

	// ExitSchema_name is called when exiting the schema_name production.
	ExitSchema_name(c *Schema_nameContext)

	// ExitNot_tag is called when exiting the not_tag production.
	ExitNot_tag(c *Not_tagContext)

	// ExitDebug_tag is called when exiting the debug_tag production.
	ExitDebug_tag(c *Debug_tagContext)

	// ExitColumn_tag is called when exiting the column_tag production.
	ExitColumn_tag(c *Column_tagContext)

	// ExitPendant_tag is called when exiting the pendant_tag production.
	ExitPendant_tag(c *Pendant_tagContext)

	// ExitUnique_tag is called when exiting the unique_tag production.
	ExitUnique_tag(c *Unique_tagContext)

	// ExitPartition_tag is called when exiting the partition_tag production.
	ExitPartition_tag(c *Partition_tagContext)

	// ExitRow_tag is called when exiting the row_tag production.
	ExitRow_tag(c *Row_tagContext)

	// ExitAs_tag is called when exiting the as_tag production.
	ExitAs_tag(c *As_tagContext)

	// ExitFrom_tag is called when exiting the from_tag production.
	ExitFrom_tag(c *From_tagContext)

	// ExitInto_tag is called when exiting the into_tag production.
	ExitInto_tag(c *Into_tagContext)

	// ExitWork_tag is called when exiting the work_tag production.
	ExitWork_tag(c *Work_tagContext)

	// ExitWith_grant_option is called when exiting the with_grant_option production.
	ExitWith_grant_option(c *With_grant_optionContext)

	// ExitWith_admin_option is called when exiting the with_admin_option production.
	ExitWith_admin_option(c *With_admin_optionContext)

	// ExitTime_zone_or_local is called when exiting the time_zone_or_local production.
	ExitTime_zone_or_local(c *Time_zone_or_localContext)

	// ExitSub_plsql_datatype is called when exiting the sub_plsql_datatype production.
	ExitSub_plsql_datatype(c *Sub_plsql_datatypeContext)

	// ExitDatatype_list is called when exiting the datatype_list production.
	ExitDatatype_list(c *Datatype_listContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)

	// ExitDatatype2 is called when exiting the datatype2 production.
	ExitDatatype2(c *Datatype2Context)

	// ExitOpr_dtype is called when exiting the opr_dtype production.
	ExitOpr_dtype(c *Opr_dtypeContext)

	// ExitOpr_datatype_lst is called when exiting the opr_datatype_lst production.
	ExitOpr_datatype_lst(c *Opr_datatype_lstContext)

	// ExitInterval_qualifier is called when exiting the interval_qualifier production.
	ExitInterval_qualifier(c *Interval_qualifierContext)

	// ExitDtype is called when exiting the dtype production.
	ExitDtype(c *DtypeContext)

	// ExitDtype1 is called when exiting the dtype1 production.
	ExitDtype1(c *Dtype1Context)

	// ExitDtype2 is called when exiting the dtype2 production.
	ExitDtype2(c *Dtype2Context)

	// ExitDouble_length_option is called when exiting the double_length_option production.
	ExitDouble_length_option(c *Double_length_optionContext)

	// ExitSize_unit_caluse is called when exiting the size_unit_caluse production.
	ExitSize_unit_caluse(c *Size_unit_caluseContext)

	// ExitLt_integer_negative is called when exiting the lt_integer_negative production.
	ExitLt_integer_negative(c *Lt_integer_negativeContext)

	// ExitCreate_contextindex_stmt is called when exiting the create_contextindex_stmt production.
	ExitCreate_contextindex_stmt(c *Create_contextindex_stmtContext)

	// ExitLexer_name is called when exiting the lexer_name production.
	ExitLexer_name(c *Lexer_nameContext)

	// ExitLexer_clause is called when exiting the lexer_clause production.
	ExitLexer_clause(c *Lexer_clauseContext)

	// ExitLexer_clause2 is called when exiting the lexer_clause2 production.
	ExitLexer_clause2(c *Lexer_clause2Context)

	// ExitSync is called when exiting the sync production.
	ExitSync(c *SyncContext)

	// ExitDrop_contextindex_stmt is called when exiting the drop_contextindex_stmt production.
	ExitDrop_contextindex_stmt(c *Drop_contextindex_stmtContext)

	// ExitAlter_contextindex_stmt is called when exiting the alter_contextindex_stmt production.
	ExitAlter_contextindex_stmt(c *Alter_contextindex_stmtContext)

	// ExitCti_sync_option is called when exiting the cti_sync_option production.
	ExitCti_sync_option(c *Cti_sync_optionContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitSizeof_type is called when exiting the sizeof_type production.
	ExitSizeof_type(c *Sizeof_typeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitArray_type is called when exiting the array_type production.
	ExitArray_type(c *Array_typeContext)

	// ExitBuiltin_types is called when exiting the builtin_types production.
	ExitBuiltin_types(c *Builtin_typesContext)

	// ExitIntegral_type is called when exiting the integral_type production.
	ExitIntegral_type(c *Integral_typeContext)

	// ExitSql_builtin_types is called when exiting the sql_builtin_types production.
	ExitSql_builtin_types(c *Sql_builtin_typesContext)

	// ExitCursor_declaration is called when exiting the cursor_declaration production.
	ExitCursor_declaration(c *Cursor_declarationContext)

	// ExitCursor_declaration_2 is called when exiting the cursor_declaration_2 production.
	ExitCursor_declaration_2(c *Cursor_declaration_2Context)

	// ExitCursor_attrs_options is called when exiting the cursor_attrs_options production.
	ExitCursor_attrs_options(c *Cursor_attrs_optionsContext)

	// ExitCursor_attrs is called when exiting the cursor_attrs production.
	ExitCursor_attrs(c *Cursor_attrsContext)

	// ExitCursor_attr is called when exiting the cursor_attr production.
	ExitCursor_attr(c *Cursor_attrContext)

	// ExitOpt_rank_specifier is called when exiting the opt_rank_specifier production.
	ExitOpt_rank_specifier(c *Opt_rank_specifierContext)

	// ExitRank_specifiers is called when exiting the rank_specifiers production.
	ExitRank_specifiers(c *Rank_specifiersContext)

	// ExitRank_specifier is called when exiting the rank_specifier production.
	ExitRank_specifier(c *Rank_specifierContext)

	// ExitOpt_dim_separators is called when exiting the opt_dim_separators production.
	ExitOpt_dim_separators(c *Opt_dim_separatorsContext)

	// ExitOpt_rank_specifier2 is called when exiting the opt_rank_specifier2 production.
	ExitOpt_rank_specifier2(c *Opt_rank_specifier2Context)

	// ExitDim_separators is called when exiting the dim_separators production.
	ExitDim_separators(c *Dim_separatorsContext)

	// ExitOpt_argument_list is called when exiting the opt_argument_list production.
	ExitOpt_argument_list(c *Opt_argument_listContext)

	// ExitJson_fun_tail is called when exiting the json_fun_tail production.
	ExitJson_fun_tail(c *Json_fun_tailContext)

	// ExitIgnore_nulls_clause is called when exiting the ignore_nulls_clause production.
	ExitIgnore_nulls_clause(c *Ignore_nulls_clauseContext)

	// ExitMixed_param_list is called when exiting the mixed_param_list production.
	ExitMixed_param_list(c *Mixed_param_listContext)

	// ExitMixed_param is called when exiting the mixed_param production.
	ExitMixed_param(c *Mixed_paramContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitCursor_option is called when exiting the cursor_option production.
	ExitCursor_option(c *Cursor_optionContext)

	// ExitWithout_into_select2 is called when exiting the without_into_select2 production.
	ExitWithout_into_select2(c *Without_into_select2Context)

	// ExitCursor_option_2 is called when exiting the cursor_option_2 production.
	ExitCursor_option_2(c *Cursor_option_2Context)

	// ExitRegion_size is called when exiting the region_size production.
	ExitRegion_size(c *Region_sizeContext)

	// ExitCopy_num is called when exiting the copy_num production.
	ExitCopy_num(c *Copy_numContext)

	// ExitRedundancy_clause is called when exiting the redundancy_clause production.
	ExitRedundancy_clause(c *Redundancy_clauseContext)

	// ExitStriping_clause is called when exiting the striping_clause production.
	ExitStriping_clause(c *Striping_clauseContext)

	// ExitWith_huge_clause is called when exiting the with_huge_clause production.
	ExitWith_huge_clause(c *With_huge_clauseContext)
}
