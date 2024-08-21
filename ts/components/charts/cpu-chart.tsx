'use client'
import React from 'react'

import { PprofType } from '@/components/charts/option/use-graph-data'
import { BasicGraph } from '@/components/charts/basic-graph'
import { usePprofOption } from '@/components/charts/option/use-pprof-option'

export const CPUGraph = () => {
  let [option, refreshKey] = usePprofOption({ pprofType: PprofType.cpu })

  return <BasicGraph option={option} refreshKey={refreshKey} />
}
